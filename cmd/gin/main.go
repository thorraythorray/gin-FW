package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx"
	"github.com/thorraythorray/go-proj/global"
	_ "github.com/thorraythorray/go-proj/initialize"
)

func main() {
	// 读取配置参数
	gin.SetMode(global.Mode)
	global.Logger.Infof("正使用%s环境,配置文件%s", global.Mode, global.Confile)

	engine := gin.Default()
	// 初始化app
	ginx.InitMoudles(engine)

	engine.Run(
		fmt.Sprintf("%s:%s", global.Config.Server.Host, global.Config.Server.Port),
	)
}
