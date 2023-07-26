package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/router"
)

func RouterRegister(R *gin.Engine) {
	apiRouter := R.Group("/v1")
	// user moudle
	router.UserRouter(apiRouter)
	router.JwtRouter(apiRouter)
}
