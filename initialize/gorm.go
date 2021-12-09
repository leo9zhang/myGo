package initialize

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"myGo/global"
	"myGo/model/customer"
	"os"
)

func Gorm() *gorm.DB {
	switch global.ServerConfig.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		//return GormPgSql()
		return nil
	default:
		return GormMysql()
	}
}

func GormMysql() *gorm.DB {
	m := global.ServerConfig.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(),
		SkipInitializeWithVersion: false,
		DefaultStringSize:         191,
		DefaultDatetimePrecision:  nil,
		DisableDatetimePrecision:  false,
		DontSupportRenameIndex:    false,
		DontSupportRenameColumn:   false,
		DontSupportForShareClause: false,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// RegisterTables 注册数据库表专用
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		//system.SysUser{},
		customer.Customer{},
	)
	if err != nil {
		global.Logger.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.Logger.Info("register table success")
}
