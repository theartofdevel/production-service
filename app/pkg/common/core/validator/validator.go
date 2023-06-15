package validator

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
	"time"
)

type Validator interface {
	Validate() error
}

var validate *validator.Validate

func New(structDateFormat string) error {
	validate = validator.New()

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		tagValue := field.Tag.Get("json")
		if tagValue == "" {
			return field.Name
		}

		fieldName := strings.SplitN(tagValue, ",", 2)[0] //nolint:gomnd
		if fieldName == "-" {
			return field.Name
		}

		return fieldName
	})

	err := validate.RegisterValidation("date", func(fl validator.FieldLevel) bool {
		val := fl.Field().String()
		if val == "" {
			return true
		}

		_, err := time.Parse(structDateFormat, fl.Field().String())
		return err == nil
	})
	if err != nil {
		return err
	}

	return nil
}
