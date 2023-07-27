package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/api"
)

func AdminRouter(r *gin.RouterGroup) {
	adminRouter := r.Group("/")

	adminRouter.POST("/token/obtain", api.AdminApiImpl.ObtainToken)

	adminRouter.GET("/user/list", api.AdminApiImpl.GetUsers)
	adminRouter.POST("/user/create", api.AdminApiImpl.CreateUser)
}
