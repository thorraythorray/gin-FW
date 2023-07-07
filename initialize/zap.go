package initialize

import (
	"github.com/thorraythorray/go-proj/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func ZapInit() {
	zapConfig := global.ConfigData.Zap

	// logger, _ := zap.NewProduction()

	encoderConfig := zap.NewProductionEncoderConfig()
	// 重新选择时间格式函数
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	logger := zap.New(
		zapcore.NewCore(
			// zapcore.NewJSONEncoder(encoderConfig),
			zapcore.NewConsoleEncoder(encoderConfig),
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
	// return logger
	global.Logger = logger
}
