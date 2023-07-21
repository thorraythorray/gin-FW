package ginx

import (
	"os"

	"github.com/thorraythorray/go-proj/ginx/schema"
	"github.com/thorraythorray/go-proj/global"
)

func MakeMigration() {
	err := global.DB.AutoMigrate(
		// admin models
		schema.User{},
	)
	if err != nil {
		global.Logger.Errorf("DB migrate err:%s", err.Error())
		os.Exit(0)
	}
}
