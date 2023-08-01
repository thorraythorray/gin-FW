package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func newValidator() *validator.Validate {
	v := validator.New()
	v.RegisterValidation("phone", PhoneReg)
	return v
}

func ValidateWithSturct(s interface{}) string {
	v := newValidator()
	tagList := []string{}
	if errs := v.Struct(s); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			tagList = append(tagList, strings.ToLower(err.Field()))
		}
	}
	errString := strings.Join(tagList, ",") + "等参数格式有误"
	return errString
}
