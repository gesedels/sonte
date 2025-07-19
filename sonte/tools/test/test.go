// Package test implements unit testing data and functions.
package test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockFiles is a map of mock JSON files for unit testing.
var MockFiles = map[string]map[string]any{
	"1000.json": {"time": "1000", "body": "First note. #test\n"},
	"2000.json": {"time": "2000", "body": "Second note. #test\n"},
	"list.json": {"test": []any{"1000", "2000"}},
}

// AssertFile asserts a file's body is equal to a JSON value.
func AssertFile(t *testing.T, orig string, want any) bool {
	bytes, err := os.ReadFile(orig)
	if err != nil {
		panic(err)
	}

	var jval any
	if err := json.Unmarshal(bytes, &jval); err != nil {
		panic(err)
	}

	return assert.Equal(t, want, jval)
}

// MockDire returns a path to a temporary directory populated with MockFiles.
func MockDire(t *testing.T) string {
	dire := t.TempDir()

	for base, jval := range MockFiles {
		bytes, err := json.Marshal(jval)
		if err != nil {
			panic(err)
		}

		dest := filepath.Join(dire, base)
		if err := os.WriteFile(dest, bytes, 0666); err != nil {
			panic(err)
		}
	}

	return dire
}

// MockFile returns a path to a temporary file populated with a MockFiles entry.
func MockFile(t *testing.T, base string) string {
	bytes, err := json.Marshal(MockFiles[base])
	if err != nil {
		panic(err)
	}

	dire := t.TempDir()
	dest := filepath.Join(dire, base)
	if err := os.WriteFile(dest, bytes, 0666); err != nil {
		panic(err)
	}

	return dest
}
