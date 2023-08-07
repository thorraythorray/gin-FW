package ginx

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/internal"
	"github.com/thorraythorray/go-proj/ginx/middleware"
	"github.com/thorraythorray/go-proj/ginx/schema"
	"github.com/thorraythorray/go-proj/global"
	"github.com/thorraythorray/go-proj/pkg/rbac"
)

func makeMigrations() {
	err := global.DB.AutoMigrate(
		// admin models
		&schema.UserModel{},
		&schema.RoleModel{},
	)
	if err != nil {
		global.Logger.Errorf("DB migrate err:%s", err.Error())
		os.Exit(0)
	}
}

func initCasbin(R *gin.Engine) {
	e := rbac.CasbinImpl(global.DB)

	// 默认添加用户所有的权限
	routes := R.Routes()
	for _, route := range routes {
		e.AddPolicy("admin", route.Path, route.Method)
	}

	// 从固定目录读取权限
	for _, v := range internal.DefaultCasbinRules {
		e.AddPolicies(v.Role, v.CasbinInfos)
	}
}

func InitMoudles(R *gin.Engine) {
	// cros->error->logger
	R.Use(
		middleware.CrosMiddleware(),
		middleware.RecoverMiddleware(),
	)
	makeMigrations()

	RouterRegister(R)
}
