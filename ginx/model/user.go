package model

type User struct {
	BaseModel
	Username string `json:"username" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Status   uint8
	Identity string `json:"identity" gorm:"not null"`
}
