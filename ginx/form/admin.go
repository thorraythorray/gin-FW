package form

type Register struct {
	Username string `json:"username" validate:"required,min=5,max=20"`
	Password string `json:"password" validate:"required,min=5,max=20"`
	Phone    string `json:"phone" validate:"phone"`
	Email    string `json:"email" validate:"email"`
}

func (form *Register) Validate() error {
	return nil
}

type Login struct {
	Username string `json:"username" validate:"required,min=5,max=20"`
	Password string `json:"password" validate:"required,min=5,max=20"`
}

func (form *Login) Validate() error {
	return nil
}
