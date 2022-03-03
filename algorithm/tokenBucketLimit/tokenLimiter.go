package tokenBucketLimit

import (
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/redis"
	xrate "golang.org/x/time/rate"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

const (
	// to be compatible with aliyun redis, we cannot use `local key = KEYS[1]` to reuse the key
	// KEYS[1] as tokens_key
	// KEYS[2] as timestamp_key
	script = `local rate = tonumber(ARGV[1])
local capacity = tonumber(ARGV[2])
local now = tonumber(ARGV[3])
local requested = tonumber(ARGV[4])
local fill_time = capacity/rate
local ttl = math.floor(fill_time*2)
local last_tokens = tonumber(redis.call("get", KEYS[1]))
if last_tokens == nil then
    last_tokens = capacity
end

local last_refreshed = tonumber(redis.call("get", KEYS[2]))
if last_refreshed == nil then
    last_refreshed = 0
end

local delta = math.max(0, now-last_refreshed)
local filled_tokens = math.min(capacity, last_tokens+(delta*rate))
local allowed = filled_tokens >= requested
local new_tokens = filled_tokens
if allowed then
    new_tokens = filled_tokens - requested
end

redis.call("setex", KEYS[1], ttl, new_tokens)
redis.call("setex", KEYS[2], ttl, now)

return allowed`
	tokenFormat     = "{%s}.tokens"
	timestampFormat = "{%s}.ts"
	pingInterval    = time.Millisecond * 100
)

type TokenLimiter struct {
	// 每秒生产速率
	rate int
	// 桶容量
	burst int
	// 存储容器
	store *redis.Redis
	// redis key
	tokenKey string
	// 桶刷新时间key
	timestampKey string
	// lock
	rescueLock sync.Mutex
	// redis健康标识
	redisAlive uint32
	// redis故障时采用进程内 令牌桶限流器
	rescueLimiter *xrate.Limiter
	// redis监控探测任务标识
	monitorStarted bool
}

func NewTokenLimiter(rate, burst int, store *redis.Redis, key string) *TokenLimiter {
	tokenKey := fmt.Sprintf(tokenFormat, key)
	timestampKey := fmt.Sprintf(timestampFormat, key)

	return &TokenLimiter{
		rate:          rate,
		burst:         burst,
		store:         store,
		tokenKey:      tokenKey,
		timestampKey:  timestampKey,
		redisAlive:    1,
		rescueLimiter: xrate.NewLimiter(xrate.Every(time.Second/time.Duration(rate)), burst),
	}
}

func (lim *TokenLimiter) reserveN(now time.Time, n int) bool {
	// 判断redis是否健康
	// redis故障时采用进程内限流器
	// 兜底保障
	if atomic.LoadUint32(&lim.redisAlive) == 0 {
		return lim.rescueLimiter.AllowN(now, n)
	}
	// 执行脚本获取令牌
	resp, err := lim.store.Eval(
		script,
		[]string{
			lim.tokenKey,
			lim.timestampKey,
		},
		[]string{
			strconv.Itoa(lim.rate),
			strconv.Itoa(lim.burst),
			strconv.FormatInt(now.Unix(), 10),
			strconv.Itoa(n),
		})
	// redis allowed == false
	// Lua boolean false -> r Nil bulk reply
	// 特殊处理key不存在的情况
	if err == redis.Nil {
		return false
	} else if err != nil {
		logx.Errorf("fail to use rate limiter: %s, use in-process limiter for rescue", err)
		// 执行异常，开启redis健康探测任务
		// 同时采用进程内限流器作为兜底
		lim.startMonitor()
		return lim.rescueLimiter.AllowN(now, n)
	}

	code, ok := resp.(int64)
	if !ok {
		logx.Errorf("fail to eval redis script: %v, use in-process limiter for rescue", resp)
		lim.startMonitor()
		return lim.rescueLimiter.AllowN(now, n)
	}

	// redis allowed == true
	// Lua boolean true -> r integer reply with value of 1
	return code == 1
}

// 开启redis健康探测
func (lim *TokenLimiter) startMonitor() {
	lim.rescueLock.Lock()
	defer lim.rescueLock.Unlock()
	// 防止重复开启
	if lim.monitorStarted {
		return
	}

	// 设置任务和健康标识
	lim.monitorStarted = true
	atomic.StoreUint32(&lim.redisAlive, 0)
	// 健康探测
	go lim.waitForRedis()
}

// redis健康探测定时任务
func (lim *TokenLimiter) waitForRedis() {
	ticker := time.NewTicker(pingInterval)
	// 健康探测成功时回调此函数
	defer func() {
		ticker.Stop()
		lim.rescueLock.Lock()
		lim.monitorStarted = false
		lim.rescueLock.Unlock()
	}()

	for range ticker.C {
		// ping属于redis内置健康探测命令
		if lim.store.Ping() {
			// 健康探测成功，设置健康标识
			atomic.StoreUint32(&lim.redisAlive, 1)
			return
		}
	}
}
