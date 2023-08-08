package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/api"
)

func UserRouter(r *gin.RouterGroup) {
	user := r.Group("/")
	user.GET("/users", api.AdminApiImpl.GetUsers)

}
