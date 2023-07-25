package model

type User struct {
	BaseModel
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Status   uint8
}
