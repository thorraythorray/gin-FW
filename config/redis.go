package config

import (
	"github.com/thorraythorray/go-proj/config/internal"
)

type RedisConf struct {
	internal.GeneralDBConfig `yaml:",inline" mapstructure:",squash"`
	DB                       int `mapstructure:"db" json:"db" yaml:"db"`
}
