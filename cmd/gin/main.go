package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/global"
	_ "github.com/thorraythorray/go-proj/initialize"
)

func main() {
	// 读取配置参数
	gin.SetMode(global.Mode)
	global.Logger.Infof("正在使用%s环境,config的路径为%s", global.Mode, global.ConfigPath)

}
