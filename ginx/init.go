package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/middleware"
)

func InitMoudles(R *gin.Engine) {
	RouterRegister(R)

	R.Use(middleware.RecoverMiddleware())

	MakeMigration()
}
