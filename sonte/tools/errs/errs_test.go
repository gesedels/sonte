package errs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// success
	err := New("%s", "test")
	assert.Equal(t, "test", err.Text)
	assert.Nil(t, err.Wrap)
}

func TestWrap(t *testing.T) {
	// setup
	wrap := errors.New("wrap")

	// success
	err := Wrap(wrap, "%s", "test")
	assert.Equal(t, "test", err.Text)
	assert.Equal(t, wrap, err.Wrap)
}

func TestError(t *testing.T) {
	// setup
	err := New("test")

	// success
	text := err.Error()
	assert.Equal(t, "test", text)
}

func TestUnwrap(t *testing.T) {
	// setup
	wrap := errors.New("wrap")
	err := Wrap(wrap, "test")

	// success
	uerr := err.Unwrap()
	assert.Equal(t, wrap, uerr)
}
