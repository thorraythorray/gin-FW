package initialize

import (
	"github.com/thorraythorray/go-proj/config"
	"github.com/thorraythorray/go-proj/global"
)

func init() {
	// 初始化运行模式
	global.Mode, global.ConfigPath = config.ModeObtain()
	Viper()
	MySQLConnect()
	RedisConnect()
	ZapConsoleInit()
}
