package schema

type SchemaHandler interface {
	Validate() error
}

func CustomValidate(f SchemaHandler) error {
	return f.Validate()
}
