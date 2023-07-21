package initialize

import (
	"flag"
	"fmt"

	"github.com/thorraythorray/go-proj/global"
	"github.com/thorraythorray/go-proj/pkg/helper"
)

func modeObtain() {
	flag.StringVar(&global.Mode, "m", "release", "choose mod in [debug release test]")
	flag.Parse()

	global.Confile = "config." + global.Mode + ".yaml"
	exist, _ := helper.PathExist(global.Confile)
	if !exist {
		panic(fmt.Errorf("fatal error: check config file"))
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
