package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type _Validator struct {
	Instance *validator.Validate
}

func (v *_Validator) Register() {
	v.Instance.RegisterValidation("phone", ValidatePhone)
}

func (v *_Validator) StructValidate(s interface{}) string {
	var tagList []string
	if errs := v.Instance.Struct(s); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			tagList = append(tagList, err.Tag())
		}
	}
	errString := strings.Join(tagList, ",") + "等参数格式有误"
	return errString
}

func Validate(s interface{}) string {
	v := _Validator{
		Instance: validator.New(),
	}
	v.Register()
	errMsg := v.StructValidate(s)
	return errMsg
}
