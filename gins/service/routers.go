package service

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/gins/service/admin"
)

func RouterInit(Router *gin.Engine) {
	apiRouter := Router.Group("/v1")
	admin.AdminRouter(apiRouter)
}
