package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"myGo/global"
)

func Redis() *redis.Client {
	redisCfg := global.ServerConfig.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Logger.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.Logger.Info("redis connect ping response:", zap.String("pong", pong))
		global.Redis = client
	}
	return client
}
