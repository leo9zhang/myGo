package core

import (
	"fmt"
	"go.uber.org/zap"
	"myGo/global"
	"myGo/initialize"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.ServerConfig.System.Addr)
	s := initServer(address, Router)

	time.Sleep(10 * time.Microsecond)
	global.Logger.Info("server run success on ", zap.String("address", address))

	global.Logger.Error(s.ListenAndServe().Error())
}
