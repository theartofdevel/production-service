package errors

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ErrCustom = New("my custom error")

func TestWrapCauseRootError(t *testing.T) {
	err1 := errors.Wrap(ErrCustom, "err1")
	err2 := errors.Wrap(err1, "err2")
	err3 := errors.Wrap(err2, "err3")

	assert.True(t, errors.Is(errors.Cause(err3), ErrCustom))
}
