package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type structValidator struct {
	s interface{}
}

func StructValidator(s interface{}) Validator {
	return &structValidator{s: s}
}

func (v *structValidator) Validate() error {
	if validate == nil {
		panic("need to call SetValidate() before this action")
	}

	if err := validate.Struct(v.s); err != nil {
		switch vErr := err.(type) {
		case *validator.InvalidValidationError:
			return err
		case validator.ValidationErrors:
			errFields := ErrorFields{}

			for _, fieldErr := range vErr {
				errFields[fieldErr.Field()] = fmt.Sprintf(
					"field validation for '%s' failed on the '%s' tag",
					fieldErr.Field(), fieldErr.Tag(),
				)
			}

			return ValidationError{Fields: errFields}
		}
	}

	return nil
}
