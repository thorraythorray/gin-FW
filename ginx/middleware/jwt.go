package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/thorraythorray/go-Jarvis/admin/auth"
	"github.com/thorraythorray/go-proj/ginx/internal"
	"github.com/thorraythorray/go-proj/global"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenstring := c.Request.Header.Get("G-Token")
		if tokenstring == "" {
			panic("required G-Token")
		}
		jwt := auth.JWT{
			SigningKey: internal.JwtSignKey,
		}
		claims, _, err := auth.AuthorizerImpl.Authenticate(&jwt, tokenstring)
		if err == nil {
			if n, ok := claims.(*auth.NewJwtClaim); ok {
				// 不使用casbin的role_definition
				// var user schema.AuthorizedUserInfo
				// err1 := global.DB.Table("users").
				// 	Select("users.id as user_id, users.username as user_name, roles.id as role_id, roles.name as role_name").
				// 	Joins("left join roles on roles.id = users.role_id").
				// 	Where("users.id = ?", n.UserIdtentify).Scan(&user).Error
				// if err == nil {
				// 	global.User = &user
				// 	c.Next()
				// } else {
				// 	panic(err1)
				// }
				global.User = n.UserIdtentify
				c.Next()
			}
		} else {
			panic(err)
		}
	}
}
