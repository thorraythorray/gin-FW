package internal

type GeneralDBConfig struct {
	Host   string `mapstructure:"host" json:"host" yaml:"host"`
	Port   string `mapstructure:"port" json:"port" yaml:"port"`
	User   string `mapstructure:"user" json:"user" yaml:"user"`
	Passwd string `mapstructure:"passwd" json:"passwd" yaml:"passwd"`
}
