package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"myGo/config"
)

var (
	Db           *gorm.DB
	Redis        *redis.Client
	ServerConfig config.Server
	Viper        *viper.Viper
	Logger       *zap.Logger
)
