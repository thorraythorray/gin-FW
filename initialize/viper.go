package initialize

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/thorraythorray/go-proj/global"
)

func Viper() {
	v := viper.New()
	v.SetConfigFile(global.Confile)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		global.Logger.Errorf("viper read config err:%s", err.Error())
		os.Exit(0)
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
