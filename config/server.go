package config

import "github.com/thorraythorray/go-proj/config/internal"

type ServerConfig struct {
	internal.GeneralConfig `yaml:",inline" mapstructure:",squash"`
	Mode                   string
}
