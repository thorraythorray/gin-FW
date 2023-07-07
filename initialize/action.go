package initialize

import (
	"github.com/thorraythorray/go-proj/config"
	"github.com/thorraythorray/go-proj/global"
)

func init() {
	// 初始化运行模式
	global.Mode, global.ConfigPath = config.ModeObtain()
	// 获取配置文件内容
	Viper()
	MySQLConnect()
	RedisConnect()
	ZapInit()
}
