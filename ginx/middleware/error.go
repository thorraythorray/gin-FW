package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/api/response"
	"github.com/thorraythorray/go-proj/global"
)

func RecoverMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.Errorf("Panic:", err)
				response.ResponseWithMsg(c, http.StatusInternalServerError, err)
				c.Abort()
			}
		}()
		c.Next()
	}
}
