package main

import (
	"myGo/core"
	"myGo/global"
	"myGo/initialize"
)

func main() {
	global.Viper = core.Viper()   // 初始化Viper
	global.Logger = core.Zap()    // 初始化zap日志库
	global.Db = initialize.Gorm() // gorm连接数据库

	global.Redis = initialize.Redis()

	if global.Db != nil {
		initialize.RegisterTables(global.Db) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.Db.DB()
		defer db.Close()
	}
	core.RunServer()
}
