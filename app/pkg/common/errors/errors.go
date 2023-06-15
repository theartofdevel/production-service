package errors

import (
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

func New(msg string) error {
	return errors.New(msg)
}
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}
func Is(err, target error) bool {
	return errors.Is(err, target)
}
func Unwrap(err error) error {
	return errors.Unwrap(err)
}
func Wrap(err error, msg string) error {
	return errors.Wrap(err, msg)
}
func Cause(err error) error {
	return errors.Cause(err)
}
func Append(err error, errs ...error) *multierror.Error {
	return multierror.Append(err, errs...)
}
func Flatten(err error) error {
	return multierror.Flatten(err)
}
func Prefix(err error, prefix string) error {
	return multierror.Prefix(err, prefix)
}
