package initialize

import (
	"time"

	"github.com/thorraythorray/go-proj/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MySQLConnect(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         64,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func MySQLPoolInit() {
	m := global.Config.Mysql
	db := MySQLConnect(m.Dsn())
	dbPool, err := db.DB()
	if err != nil {
		panic(err)
	}
	// 暂时用默认
	// dbPool.SetMaxIdleConns(10)
	// dbPool.SetMaxOpenConns(100)
	dbPool.SetConnMaxLifetime(time.Hour)
	global.DBPool = dbPool
}
