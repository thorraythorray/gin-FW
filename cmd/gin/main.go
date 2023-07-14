package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/global"
	_ "github.com/thorraythorray/go-proj/initialize"
	"github.com/thorraythorray/go-proj/pkg/gin/middleware"
	"github.com/thorraythorray/go-proj/service"
)

func main() {
	var (
		logger     = global.Logger
		config     = global.ConfigData
		mod        = global.Mode
		configPath = global.ConfigPath
	)

	// 读取配置参数
	gin.SetMode(mod)
	logger.Infof("正在使用%s环境,配置文件%s", mod, configPath)
	engine := gin.Default()
	// register router
	service.RouterInit(engine)
	// use middleware
	engine.Use(middleware.RecoverMiddleware())
	// engine.Use(middleware.LoggerRequestMiddleware())
	serverAddr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	engine.Run(serverAddr)
}
