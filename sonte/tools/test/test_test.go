package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertDire(t *testing.T) {
	// setup
	dire := t.TempDir()
	dest := filepath.Join(dire, "name.extn")
	os.WriteFile(dest, []byte("Body.\n"), 0666)

	// success
	AssertDire(t, dire, map[string]string{"name.extn": "Body.\n"})
}

func TestAssertFile(t *testing.T) {
	// setup
	dire := t.TempDir()
	dest := filepath.Join(dire, "name.extn")
	os.WriteFile(dest, []byte("Body.\n"), 0666)

	// success
	AssertFile(t, dest, "Body.\n")
}

func TestMockDire(t *testing.T) {
	// success
	dire := MockDire(t)
	assert.DirExists(t, dire)

	for base, body := range MockData {
		orig := filepath.Join(dire, base)
		bytes, _ := os.ReadFile(orig)
		assert.Equal(t, body, string(bytes))
	}
}

func TestMockFile(t *testing.T) {
	// success
	orig := MockFile(t, "alpha.extn")
	bytes, _ := os.ReadFile(orig)
	assert.Equal(t, MockData["alpha.extn"], string(bytes))
}
