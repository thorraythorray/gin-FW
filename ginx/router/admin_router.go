package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/api"
)

func AdminRouter(r *gin.RouterGroup) {
	adminRouter := r.Group("/")

	adminRouter.POST("/login", api.AdminApiImpl.Login)
	adminRouter.GET("/users", api.AdminApiImpl.GetUsers)
	adminRouter.POST("/user/register", api.AdminApiImpl.Register)
}