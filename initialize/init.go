package initialize

import (
	"flag"
	"os"

	"github.com/thorraythorray/go-proj/global"
	"github.com/thorraythorray/go-proj/pkg/helper"
)

func modeObtain() {
	flag.StringVar(&global.Mode, "m", "release", "choose mod in [debug release test]")
	flag.Parse()

	global.Confile = "config." + global.Mode + ".yaml"
	exist, _ := helper.PathExist(global.Confile)
	if !exist {
		global.Logger.Errorf("%s config file not exist", global.Confile)
		os.Exit(0)
	}
}

func init() {
	// 初始化运行模式
	modeObtain()
	Viper()
	MySQLPoolInit()
	RedisInit()
	ZapConsoleInit()
}
