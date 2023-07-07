package config

import (
	"flag"
	"fmt"

	"github.com/thorraythorray/go-proj/config/internal"
)

type ConfigMap struct {
	Mysql MySQLConf `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis RedisConf `mapstructure:"redis" json:"redis" yaml:"redis"`
	Zap   ZapConfig `mapstructure:"zap" json:"zap" yaml:"zap"`
	Cros CrosConfig `mapstructure:"cros" json:"cros" yaml:"cros"`
}

// 定义对db配置文件获取Data Source Name的接口
type DSN interface {
	Dsn() string
}

func ModeObtain() (string, string) {
	var mode, cfgPath string

	flag.StringVar(&mode, "m", "release", "choose mod in [debug release test]")
	flag.Parse()

	switch mode {
	case "debug":
		cfgPath = internal.ConfigDebugFile
	case "release":
		cfgPath = internal.ConfigReleaseFile
	case "test":
		cfgPath = internal.ConfigTestFile
	default:
		panic(fmt.Errorf("mod not support"))
		// todo: raise error
	}
	return mode, cfgPath
}
