package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/api"
)

func PublicRouter(r *gin.RouterGroup) {
	public := r.Group("/")
	public.POST("/login", api.AdminApiImpl.Login)
	public.POST("/register", api.AdminApiImpl.Register)
	public.POST("/role", api.AdminApiImpl.CreateRole)
	public.DELETE("/user/:id", api.AdminApiImpl.DeleteUser)
}
