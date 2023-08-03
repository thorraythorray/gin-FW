package ginx

import (
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
}
