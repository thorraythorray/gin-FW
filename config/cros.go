package config

type CrosConfig struct {
	Mode      string   `mapstructure:"database" json:"database" yaml:"database"`
	Whitelist []string `mapstructure:"whitelist" json:"whitelist" yaml:"whitelist"`
}
