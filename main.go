package main

import (
	"myGo/core"
	"myGo/global"
	"myGo/initialize"
)

func main() {
	global.GVA_VP = core.Viper()      // 初始化Viper
	global.GVA_LOG = core.Zap()       // 初始化zap日志库
	global.GVA_DB = initialize.Gorm() // gorm连接数据库

	global.GVA_REDIS = initialize.Redis()

	if global.GVA_DB != nil {
		//initialize.RegisterTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.RunServer()
}
