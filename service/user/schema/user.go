package schema

import (
	"github.com/thorraythorray/go-proj/pkg/gorm"
)

type User struct {
	gorm.BaseGorm
	Username string `form:"username" binding:"required" json:"username"`
	Password string `form:"password" binding:"required" json:"password"`
	Phone    int    `form:"phone" json:"phone"`
	Email    string `form:"email" json:"email"`
	Status   int
}
