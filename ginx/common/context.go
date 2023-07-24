package common

import "github.com/thorraythorray/go-proj/ginx/schema"

type context struct{}

func (cxt *context) RequestValidate(f schema.Former) (bool, error) {
	return f.Validate()
}

var Ctx = new(context)
