package router

import (
	"github.com/gin-gonic/gin"
)

func RouterInit(R *gin.Engine) {
	apiRouter := R.Group("/v1")

	UserRouter(apiRouter)
}
