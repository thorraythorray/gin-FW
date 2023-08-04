package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/thorraythorray/go-proj/ginx/internal"
	"github.com/thorraythorray/go-proj/global"
	"github.com/thorraythorray/go-proj/pkg/auth"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenstring := c.Request.Header.Get("G-Token")
		if tokenstring == "" {
			c.AbortWithError(http.StatusBadRequest, errors.New("缺少GToken认证参数"))
		}
		jwt := auth.JWT{
			SigningKey: internal.JwtSignKey,
		}
		claims, status, err := auth.AuthorizerImpl.Authenticate(&jwt, tokenstring)
		if err == nil {
			if n, ok := claims.(*auth.NewJwtClaim); ok {
				err := global.DB.Table("users").
					Select("users.id as user_id, users.username as user_name, role.id as role_id, role.name as role.name").
					Joins("left join roles on roles.id = users.role_id").
					Where("users.id = ?", n.UserIdtentify).Scan(global.User).Error
				if err != nil {
					c.AbortWithError(http.StatusBadRequest, err)
				}
				c.Next()
			}
		} else {
			c.AbortWithError(status, err)
		}
	}
}
