package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/internal"
	"github.com/thorraythorray/go-proj/ginx/middleware"
	"github.com/thorraythorray/go-proj/global"
	"github.com/thorraythorray/go-proj/pkg/rbac"
)

func InitMoudles(R *gin.Engine) {
	// cros->error->logger
	R.Use(
		middleware.CrosMiddleware(),
		middleware.RecoverMiddleware(),
	)
	RouterRegister(R)
	MakeMigration()

	e := rbac.CasbinImpl(global.DB)
	for _, v := range internal.DefaultCasbinRules {
		e.AddPolicies(v.Role, v.CasbinInfos)
	}

}
