package initialize

import (
	"github.com/thorraythorray/go-proj/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MySQLConnect() {
	m := global.ConfigData.Mysql
	db, err := gorm.Open(mysql.Open(m.Dsn()), &gorm.Config{})
	if err == nil {
		global.DB = db
	} else {
		panic(err)
	}
}
