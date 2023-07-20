package config

type ConfigMap struct {
	Mysql  MySQLConf    `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis  RedisConf    `mapstructure:"redis" json:"redis" yaml:"redis"`
	Zap    ZapConfig    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Cros   CrosConfig   `mapstructure:"cros" json:"cros" yaml:"cros"`
	Server ServerConfig `mapstructure:"server" json:"server" yaml:"server"`
}
