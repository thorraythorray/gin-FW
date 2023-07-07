package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()

	message := "This is a log message"

	sugar.Info(message)
	sugar.Error(message)
	sugar.Warn(message)
	sugar.Debug(message)
}
