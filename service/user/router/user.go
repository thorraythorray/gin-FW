package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/service/user/api"
)

// type RouterGroup struct{}

func UserRouter(r *gin.RouterGroup) {
	userRouter := r.Group("/user")
	userApi := new(api.UserApi)
	{
		userRouter.GET("/list", userApi.GetUsers)
	}
}
