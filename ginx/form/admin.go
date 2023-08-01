package form

type Register struct {
	Username string `json:"username" validate:"required,len=5,20"`
	Password string `json:"password" validate:"required,len=8,20"`
	Phone    string `json:"phone" validate:"phone"`
	Email    string `json:"email" validate:"email"`
}

func (form *Register) Validate() error {
	return nil
}

type Login struct {
	Username string `json:"username" validate:"required,len=5,20"`
	Password string `json:"password" validate:"required,len=8,20"`
}

func (form *Login) Validate() error {
	return nil
}
