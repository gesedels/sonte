package test

import (
	"path/filepath"
	"testing"
)

func TestAssertFile(t *testing.T) {
	// setup
	orig := MockFile(t, "1000.json")

	// success
	AssertFile(t, orig, MockFiles["1000.json"])
}

func TestMockDire(t *testing.T) {
	// success
	dire := MockDire(t)

	for base, jval := range MockFiles {
		orig := filepath.Join(dire, base)
		AssertFile(t, orig, jval)
	}
}

func TestMockFile(t *testing.T) {
	// success
	orig := MockFile(t, "1000.json")
	AssertFile(t, orig, MockFiles["1000.json"])
}
