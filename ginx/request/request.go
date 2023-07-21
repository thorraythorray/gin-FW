package request

type Former interface {
	Validate(interface{}) (bool, error)
}

func Validate(f Former, data interface{}) (bool, error) {
	return f.Validate(data)
}
