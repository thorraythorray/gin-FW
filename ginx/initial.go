package ginx

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/middleware"
)

func InitMoudles(R *gin.Engine) {
	// cros->error->logger
	R.Use(
		middleware.CrosMiddleware(),
		middleware.RecoverMiddleware(),
	)
	RouterRegister(R)
	MakeMigration()

	// e := rbac.CasbinImpl(global.DB)
	// for _, v := range internal.DefaultCasbinRules {
	// 	e.AddPolicies(v.Role, v.CasbinInfos)
	// }
	// 获取所有的路由信息
	routes := R.Routes()
	for _, route := range routes {
		fmt.Printf("Path: %s, Method: %s\n", route.Path, route.Method)
	}

}
