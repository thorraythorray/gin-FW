package request

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/api/response"
	"github.com/thorraythorray/go-proj/ginx/schema"
	"github.com/thorraythorray/go-proj/pkg/validator"
)

func Validate(c *gin.Context, req interface{}) bool {
	ok := true
	var err error
	if c.Request.Method == "GET" || c.Request.Method == "HEAD" {
		err = c.ShouldBindQuery(req)
	} else {
		err = c.ShouldBindJSON(req)
	}
	if err != nil || req == nil {
		response.RequestFailed(c, err)
		ok = false
	}
	if req != nil {
		errMsg := validator.ValidateWithSturct(req)
		if errMsg != "" {
			response.RequestFailed(c, errors.New(errMsg))
			ok = false
		}
		if f, ok := req.(schema.SchemaHandler); ok {
			err1 := schema.CustomValidate(f)
			if err != nil {
				response.RequestFailed(c, err1)
				ok = false
			}
		}
	}
	return ok
}
