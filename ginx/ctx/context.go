package ctx

import "github.com/thorraythorray/go-proj/ginx/schema"

type Context struct{}

func (cxt *Context) RequestValidate(f schema.Former) (bool, error) {
	return f.Validate()
}
