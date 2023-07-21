package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/api"
)

func AdminRouter(r *gin.RouterGroup) {
	userRouter := r.Group("/user")
	{
		userRouter.GET("/list", api.UserApiImpl.GetUsers)
	}
}
