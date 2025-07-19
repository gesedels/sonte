package test

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAssertFile(t *testing.T) {
	// setup
	dest := filepath.Join(t.TempDir(), "test.txt")
	os.WriteFile(dest, []byte("Test.\n"), 0666)

	// success
	AssertFile(t, dest, "Test.\n")
}
