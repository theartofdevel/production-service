package dal

import (
	"production_service/pkg/errors"
)

var ErrNotFound = errors.New("not found")
var ErrUniqueViolation = errors.New("unique violation")
