package util

import (
	"os"

	"github.com/thorraythorray/go-proj/ginx/model"
	"github.com/thorraythorray/go-proj/global"
)

func MakeMigration() {
	err := global.DB.AutoMigrate(
		// admin models
		model.User{},
	)
	if err != nil {
		global.Logger.Errorf("DB migrate err:%s", err.Error())
		os.Exit(0)
	}
}
