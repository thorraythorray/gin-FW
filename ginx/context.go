package ginx

import "github.com/thorraythorray/go-proj/ginx/request"

type Context struct{}

func (cxt *Context) RequestValidate(f request.Former) (bool, error) {
	return f.Validate()
}
