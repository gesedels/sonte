// Package test implements unit testing data and functions.
package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertFile asserts a file's body is equal to a string.
func AssertFile(t *testing.T, orig, body string) {
	bytes, err := os.ReadFile(orig)
	assert.Equal(t, body, string(bytes))
	assert.NoError(t, err)
}
