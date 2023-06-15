package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	for i := 0; i < 1_000; i++ {
		str, err := String(i)
		assert.NoError(t, err)
		assert.Equal(t, i, len(str))
	}
}

func BenchmarkRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := String(100)
		assert.NoError(b, err)
	}
}

func TestRandomBytes(t *testing.T) {
	for i := 0; i < 1_000; i++ {
		str, err := Bytes(i)
		assert.NoError(t, err)
		assert.Equal(t, i, len(str))
	}
}

func BenchmarkRandomBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Bytes(100)
		assert.NoError(b, err)
	}
}

func TestRandomInt(t *testing.T) {
	for i := 0; i < 1000; i++ {
		str, err := Bytes(i)
		assert.NoError(t, err)
		assert.Equal(t, i, len(str))
	}
}

func BenchmarkRandomInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Bytes(100)
		assert.NoError(b, err)
	}
}
