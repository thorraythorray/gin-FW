package initialize

import (
	"os"
	"time"

	"github.com/thorraythorray/go-proj/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func newZapEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout(time.DateTime),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func ZapFileInit() {
	var zapConfig = global.Config.Zap
	logger := zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(newZapEncoderConfig()),
			zapcore.AddSync(&lumberjack.Logger{
				Filename:   zapConfig.LogFile,
				MaxSize:    50,  // 指定日志文件大小的阈值，单位为 MB
				MaxBackups: 6,   // 最大保留的日志文件数量
				MaxAge:     180, // 日志文件保留的最大天数
				LocalTime:  true,
				Compress:   true, // 是否压缩日志文件
			}),
			zapConfig.MatchLevel(), // 设置日志级别
		),
		zap.AddCaller(),
	).Sugar()
	global.Logger = logger
}

func ZapConsoleInit() {
	var zapConfig = global.Config.Zap

	logger := zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(newZapEncoderConfig()),
			zapcore.Lock(os.Stdout),
			zapConfig.MatchLevel(), // 设置日志级别
		),
		zap.AddCaller(),
	).Sugar()
	global.Logger = logger
}
