package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/middleware"
	"github.com/thorraythorray/go-proj/ginx/router"
)

func RouterRegister(R *gin.Engine) {
	publicRouterGroup := R.Group("/api")
	{
		router.PublicRouter(publicRouterGroup)
	}

	privateRouterGroup := publicRouterGroup
	privateRouterGroup.Use(middleware.JwtAuthMiddleware())
	{
		router.UserRouter(privateRouterGroup)
	}

}
