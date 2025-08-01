package clui

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	// setup
	os.Setenv("NAME", "\tdata\n")

	// success
	data, err := Env("NAME")
	assert.Equal(t, "data", data)
	assert.NoError(t, err)

	// error - variable does not exist
	data, err = Env("NOPE")
	assert.Empty(t, data)
	assert.EqualError(t, err, `environment variable "NOPE" does not exist`)
}

func TestExec(t *testing.T) {
	// setup
	dire := t.TempDir()
	dest := filepath.Join(dire, "exec.extn")

	// success
	err := Exec("touch", dest)
	assert.FileExists(t, dest)
	assert.NoError(t, err)
}
