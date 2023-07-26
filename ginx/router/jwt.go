package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/api"
)

func JwtRouter(r *gin.RouterGroup) {
	jwtRouter := r.Group("/token")
	{
		jwtRouter.POST("/", api.ObtainToken)
	}
}
