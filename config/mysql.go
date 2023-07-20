package config

import (
	"github.com/thorraythorray/go-proj/config/internal"
)

type MySQLConf struct {
	internal.GeneralConfig `yaml:",inline" mapstructure:",squash"`
	User                   string `mapstructure:"user" json:"user" yaml:"user"`
	Passwd                 string `mapstructure:"passwd" json:"passwd" yaml:"passwd"`
	Database               string `mapstructure:"database" json:"database" yaml:"database"`
}

func (m *MySQLConf) Dsn() string {
	return m.User + ":" + m.Passwd + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
}
