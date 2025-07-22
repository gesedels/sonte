package clui

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExec(t *testing.T) {
	// setup
	dire := t.TempDir()
	dest := filepath.Join(dire, "exec.extn")

	// success
	err := Exec("touch", dest)
	assert.FileExists(t, dest)
	assert.NoError(t, err)
}

func TestEnv(t *testing.T) {
	// setup
	os.Setenv("NAME", "\tdata\n")
	os.Setenv("EMPTY", "\n")

	// success
	data, err := Env("NAME")
	assert.Equal(t, "data", data)
	assert.NoError(t, err)

	// error - variable is not set
	data, err = Env("NOPE")
	assert.Empty(t, data)
	assert.EqualError(t, err, `environment variable "NOPE" is not set`)

	// error - variable is empty
	data, err = Env("EMPTY")
	assert.Empty(t, data)
	assert.EqualError(t, err, `environment variable "EMPTY" is empty`)
}

func TestEnvs(t *testing.T) {
	// setup
	os.Setenv("NAME", "\tdata\n")
	os.Setenv("EMPTY", "\n")

	// success
	data, err := Envs("NOPE", "EMPTY", "NAME")
	assert.Equal(t, "data", data)
	assert.NoError(t, err)

	// error - variables not set or empty
	data, err = Envs("NOPE", "EMPTY")
	assert.Empty(t, data)
	assert.EqualError(t, err, `environment variables "NOPE", "EMPTY" are not set or empty`)
}
