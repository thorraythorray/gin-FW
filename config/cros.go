package config

type CrosConfig struct {
	Mode      string   `mapstructure:"mode" json:"mode" yaml:"mode"`
	Whitelist []string `mapstructure:"whitelist" json:"whitelist" yaml:"whitelist"`
}
