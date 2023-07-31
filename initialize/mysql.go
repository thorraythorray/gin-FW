package initialize

import (
	"os"
	"time"

	"github.com/thorraythorray/go-proj/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MySQLConnect(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         64,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		global.Logger.Errorf("mysql connect err:%s", err.Error())
		os.Exit(0)
	}
	return db
}

func MySQLPoolInit() {
	m := global.Config.Mysql
	db := MySQLConnect(m.Dsn())
	dbPool, err := db.DB()
	if err != nil {
		global.Logger.Errorf("mysql pool setting err:%s", err.Error())
		os.Exit(0)
	}
	// 配置连接池参数
	// dbPool.SetMaxIdleConns(10)
	// dbPool.SetMaxOpenConns(100)
	dbPool.SetConnMaxLifetime(time.Hour)
	global.DB = db
}
