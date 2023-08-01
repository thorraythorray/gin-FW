package request

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/api/response"
	"github.com/thorraythorray/go-proj/ginx/form"
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
	if err != nil {
		response.RequestFailed(c, err.Error())
		ok = false
	}
	if req != nil {
		errMsg := validator.ValidateWithSturct(req)
		if errMsg != "" {
			response.RequestFailed(c, errMsg)
			ok = false
		}
		if f, ok := req.(form.FormHandler); ok {
			err1 := form.CustomValidate(f)
			if err != nil {
				response.RequestFailed(c, err1.Error())
				ok = false
			}
		}
	}
	return ok
}
