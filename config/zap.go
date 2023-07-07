package config

import "go.uber.org/zap/zapcore"

type ZapConfig struct {
	Level   string `mapstructure:"level" json:"level" yaml:"level"`
	LogFile string `mapstructure:"logfile" json:"logfile" yaml:"logfile"`
	Format  string `mapstructure:"format" json:"format" yaml:"format"`
}

func (z *ZapConfig) MatchLevel() zapcore.Level {
	switch z.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
