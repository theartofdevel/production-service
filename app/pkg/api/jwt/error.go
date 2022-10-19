package jwt

import "github.com/pkg/errors"

var ErrBadToken = errors.New("malformed token")
