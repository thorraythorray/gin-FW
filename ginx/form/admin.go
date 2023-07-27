package form

type Register struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Phone    string `json:"phone" validate:"phone"`
	Email    string `json:"email" validate:"email"`
}
