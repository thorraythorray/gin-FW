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
				if e, ok := err.(error); ok {
					response.ResponseWithMsg(c, http.StatusBadRequest, e.Error())
				} else if ee, ok := err.(string); ok {
					response.ResponseWithMsg(c, http.StatusBadRequest, ee)
				} else {
					response.ResponseWithMsg(c, http.StatusBadRequest, e)
				}
				c.Abort()
			}
		}()
		c.Next()
	}
}
