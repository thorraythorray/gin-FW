package router

import (
	"github.com/gin-gonic/gin"
)

func RouterInit(Router *gin.Engine) {
	apiRouter := Router.Group("/v1")
	AdminRouter(apiRouter)
}
