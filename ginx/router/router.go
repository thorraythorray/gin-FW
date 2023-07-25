package router

import (
	"github.com/gin-gonic/gin"
)

func RouterRegister(R *gin.Engine) {
	apiRouter := R.Group("/v1")
	// user moudle
	UserRouter(apiRouter)
}
