package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/thorraythorray/go-proj/ginx/internal"
	"github.com/thorraythorray/go-proj/pkg/auth"
)

var jwt = auth.JWT{SigningKey: internal.SignKey}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqUser := c.Request.Header.Get("X-User")
		tokenstring := c.Request.Header.Get("X-Token")
		res, status, err := jwt.ParseJwtToken(tokenstring)
		if err == nil {
			if reqUser == res.UserID {
				c.Next()
			}
		}
		c.AbortWithError(status, err)
	}
}
