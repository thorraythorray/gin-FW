package config

import (
	"github.com/thorraythorray/go-proj/config/internal"
)

type RedisConf struct {
	internal.GeneralConfig `yaml:",inline" mapstructure:",squash"`
	Passwd                 string `mapstructure:"passwd" json:"passwd" yaml:"passwd"`
	DB                     int    `mapstructure:"db" json:"db" yaml:"db"`
}
