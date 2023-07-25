package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/middleware"
	"github.com/thorraythorray/go-proj/ginx/router"
	"github.com/thorraythorray/go-proj/ginx/util"
	"github.com/thorraythorray/go-proj/global"
	_ "github.com/thorraythorray/go-proj/initialize"
)

func main() {
	// 读取配置参数
	gin.SetMode(global.Mode)
	global.Logger.Infof("正使用%s环境,配置文件%s", global.Mode, global.Confile)

	engine := gin.Default()
	// register middleware
	engine.Use(middleware.RecoverMiddleware())
	// register router
	router.RouterRegister(engine)
	// migrate
	util.MakeMigration()

	engine.Run(
		fmt.Sprintf("%s:%s", global.Config.Server.Host, global.Config.Server.Port),
	)

}
