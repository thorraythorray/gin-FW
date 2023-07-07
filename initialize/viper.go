package initialize

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/thorraythorray/go-proj/global"
	"github.com/thorraythorray/go-proj/pkg/util"
)

func Viper() {
	cfgPath := global.ConfigPath
	// check config file exist
	exist, _ := util.PathExist(cfgPath)
	if !exist {
		panic(fmt.Errorf("fatal error check config file"))
	}

	v := viper.New()
	v.SetConfigFile(cfgPath)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.ConfigData); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.ConfigData); err != nil {
		fmt.Println(err)
	}
}
