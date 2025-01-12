package models

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
	"sync"
	"videoverse/pkg/logbox"
)

var instance *validator.Validate = nil
var once sync.Once

func NewValidator() *validator.Validate {
	once.Do(func() {
		var err error
		vrr, ok := binding.Validator.Engine().(*validator.Validate)
		if !ok {
			logbox.NewLogBox().Fatal().Err(err).Msg("failed to get validator engine")
		}

		vrr.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		instance = vrr
	})
	return instance
}

func Validate(r interface{}, errs map[string]any) map[string]any {
	var validate = NewValidator()
	err := validate.Struct(r)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errs[err.Namespace()] = err.Error()
		}
	}
	return errs
}
