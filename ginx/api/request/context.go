package request

type Former interface {
	Validate() error
}

type context struct{}

func (cxt *context) RequestValidate(f Former) error {
	return f.Validate()
}

var Ctx = new(context)
