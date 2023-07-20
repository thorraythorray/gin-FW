package initialize

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/thorraythorray/go-proj/global"
	"github.com/thorraythorray/go-proj/pkg/helper"
)

func Viper() {
	cfgFile := global.ConfFile
	// check config file exist
	exist, _ := helper.PathExist(cfgFile)
	if !exist {
		panic(fmt.Errorf("fatal error check config file"))
	}

	v := viper.New()
	v.SetConfigFile(cfgFile)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
	}
}
