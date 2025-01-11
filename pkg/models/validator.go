package models

import "github.com/go-playground/validator/v10"

var validate = validator.New(validator.WithRequiredStructEnabled())

func Validate(r interface{}, errs map[string]any) map[string]any {
	err := validate.Struct(r)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errs[err.Namespace()] = err.Error()
		}
	}
	return errs
}
