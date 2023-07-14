package service

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/service/user/router"
)

func RouterInit(Router *gin.Engine) {
	apiRouter := Router.Group("/v1")
	router.UserRouter(apiRouter)
}
