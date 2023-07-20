package schema

import "github.com/thorraythorray/go-proj/pkg/gorm/base"

type User struct {
	base.BaseModel
	Username string `form:"username" binding:"required" json:"username"`
	Password string `form:"password" binding:"required" json:"password"`
	Phone    string `form:"phone" json:"phone"`
	Email    string `form:"email" json:"email"`
	Status   uint8
}
