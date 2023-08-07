package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/global"
	"github.com/thorraythorray/go-proj/pkg/rbac"
)

func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.User == nil {
			panic("user has no role")
		}
		sub := global.User.RoleName
		obj := c.Request.URL.Path
		act := c.Request.Method
		e := rbac.CasbinImpl(global.DB).NewCasbin() // 判断策略中是否存在
		ok, _ := e.Enforce(sub, obj, act)
		if ok {
			c.Next()
		} else {
			panic("Permission Denied")
		}
	}
}
