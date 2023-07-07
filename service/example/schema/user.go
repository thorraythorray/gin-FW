package schema

import (
	"github.com/thorraythorray/go-proj/service/internal"
)

type User struct {
	internal.BaseModel
	Username string ``
	Password string
	Phone    int
	email    string
	status   int
}

// func (db *gorm.DB) CheckExists(opts ...User) {

// }
