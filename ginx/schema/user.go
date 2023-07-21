package schema

import "github.com/thorraythorray/go-proj/pkg/gorm/base"

type User struct {
	base.BaseModel
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Status   uint8
}
