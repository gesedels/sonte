// Package test implements unit testing data and functions.
package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockData is a base:body map of mock note files for unit testing.
var MockData = map[string]string{
	"alpha.extn":    "Alpha note.\n",
	"bravo.extn":    "Bravo note.\n",
	"charlie.trash": "Charlie note (trash).\n",
}

// AssertDire asserts a directory's files are equal to a base:body map.
func AssertDire(t *testing.T, dire string, files map[string]string) bool {
	for base, body := range files {
		orig := filepath.Join(dire, base)
		if !AssertFile(t, orig, body) {
			return false
		}
	}

	return true
}

// AssertFile asserts a file's body is equal to a string.
func AssertFile(t *testing.T, orig, body string) bool {
	bytes, err := os.ReadFile(orig)
	if err != nil {
		panic(err)
	}

	return assert.Equal(t, body, string(bytes))
}

// MockDire returns a temporary directory containing all MockData entries.
func MockDire(t *testing.T) string {
	dire := t.TempDir()

	for base, body := range MockData {
		dest := filepath.Join(dire, base)
		if err := os.WriteFile(dest, []byte(body), 0666); err != nil {
			panic(err)
		}
	}

	return dire
}

// MockFile returns a temporary file containing a MockData entry.
func MockFile(t *testing.T, base string) string {
	dire := t.TempDir()
	dest := filepath.Join(dire, base)
	if err := os.WriteFile(dest, []byte(MockData[base]), 0666); err != nil {
		panic(err)
	}

	return dest
}
