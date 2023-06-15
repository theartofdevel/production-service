package validator

import "fmt"

type ErrorFields map[string]string

type ValidationError struct {
	Fields ErrorFields
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%v", e.Fields)
}
