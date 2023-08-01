package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/thorraythorray/go-proj/ginx/internal"
	"github.com/thorraythorray/go-proj/pkg/auth"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqUrlPath := c.Request.URL.Path
		if !strings.Contains(internal.JwtExemptRouterString, reqUrlPath) {
			tokenstring := c.Request.Header.Get("GToken")
			if tokenstring == "" {
				c.AbortWithError(http.StatusBadRequest, errors.New("缺少GToken认证参数"))
			}
			jwt := auth.JWT{
				SigningKey: internal.JwtSignKey,
			}
			status, err := auth.AuthorizerImpl.Authenticate(&jwt, tokenstring)
			if err != nil {
				c.AbortWithError(status, err)
			}
		}
		c.Next()
	}
}
