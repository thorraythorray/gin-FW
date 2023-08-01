package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/middleware"
)

func InitMoudles(R *gin.Engine) {
	// cros->error->logger->auth
	R.Use(
		middleware.CrosMiddleware(),
		middleware.RecoverMiddleware(),
		// middleware.LoggerRequestMiddleware(),
		middleware.JwtAuthMiddleware(),
	)
	RouterRegister(R)
	MakeMigration()
}
