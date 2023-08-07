package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/global"
)

func allowCros() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, GToken")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if c.Request.Method == "OPTIONS" {
			panic("Method not allowed")
		} else {
			c.Next()
		}

	}
}

func rulesCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		whiteList := global.Config.Cros.Whitelist
		_checkInWhitelist := func(ori string, wl []string) bool {
			for _, item := range wl {
				if ori == item {
					return true
				}
			}
			return false
		}
		origin := c.Request.Header.Get("Origin")
		checkFlag := _checkInWhitelist(origin, whiteList)
		if checkFlag {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PUT")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, GToken")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Next()
		} else {
			panic("Cros not allowed")
		}
	}
}

func CrosMiddleware() gin.HandlerFunc {
	if global.Config.Cros.Mode == "allow-all" {
		return allowCros()
	} else {
		return rulesCors()
	}
}
