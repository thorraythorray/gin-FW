package initialize

import (
	"github.com/thorraythorray/go-proj/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MySQLInit() (*gorm.DB, error) {
	m := global.Config.Mysql
	db, err := gorm.Open(
		mysql.Open(m.Dsn()),
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		},
	)
	return db, err
}
