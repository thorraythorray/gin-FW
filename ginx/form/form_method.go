package form

type FormHandler interface {
	Validate() error
}

func CustomValidate(f FormHandler) error {
	return f.Validate()
}
