package schema

import (
	"os"

	"github.com/thorraythorray/go-proj/global"
)

func ModelMigrate() {
	err := global.DB.AutoMigrate(
		// admin models
		User{},
	)
	if err != nil {
		global.Logger.Errorf("DB migrate err:%s", err.Error())
		os.Exit(0)
	}
}
