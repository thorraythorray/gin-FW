package request

import "github.com/thorraythorray/go-proj/ginx/form"

type context struct{}

func (cxt *context) ValidateForm(f form.FormMethod) error {
	return f.Validate()
}

var Ctx = new(context)
