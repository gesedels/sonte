// Package test implements unit testing data and functions.
package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockNotes is a base:body map of mock notes for unit testing.
var MockNotes = map[string]string{
	"alpha.extn":    ".date 2001-02-03\n.type test\nAlpha note.\n",
	"bravo.extn":    "Bravo note.\n",
	"charlie.trash": "Charlie note (trash).\n",
}

// AssertDire asserts a directory's files are equal to a base:body map.
func AssertDire(t *testing.T, dire string, files map[string]string) {
	for base, body := range files {
		orig := filepath.Join(dire, base)
		AssertFile(t, orig, body)
	}
}

// AssertFile asserts a file's body is equal to a string.
func AssertFile(t *testing.T, orig, body string) {
	bytes, err := os.ReadFile(orig)
	assert.Equal(t, body, string(bytes))
	assert.NoError(t, err)
}

// MockDire returns a temporary directory containing all MockNotes entries.
func MockDire(t *testing.T) string {
	dire := t.TempDir()
	for base, body := range MockNotes {
		dest := filepath.Join(dire, base)
		if err := os.WriteFile(dest, []byte(body), 0666); err != nil {
			panic(err)
		}
	}

	return dire
}

// MockFile returns a temporary file containing a MockFiles entry.
func MockFile(t *testing.T, base string) string {
	dire := t.TempDir()
	dest := filepath.Join(dire, base)
	if err := os.WriteFile(dest, []byte(MockNotes[base]), 0666); err != nil {
		panic(err)
	}

	return dest
}
