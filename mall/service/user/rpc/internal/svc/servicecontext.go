package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	model2 "mall/service/user/model"
	config2 "mall/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config2.Config

	UserModel model2.UserModel
}

func NewServiceContext(c config2.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model2.NewUserModel(conn, c.CacheRedis),
	}
}
