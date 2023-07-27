package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/middleware"
)

func InitMoudles(R *gin.Engine) {
	R.Use(
		// middleware.RecoverMiddleware(),
		middleware.JwtAuthMiddleware(),
	)
	RouterRegister(R)
	MakeMigration()
}
