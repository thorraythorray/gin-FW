package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/api"
)

func UserRouter(r *gin.RouterGroup) {
	userRouter := r.Group("/user")
	{
		userRouter.GET("/list", api.UserApiImpl.GetUsers)
		userRouter.POST("/", api.UserApiImpl.CreateUser)
	}
}
