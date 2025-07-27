package errs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvarDoesNotExist(t *testing.T) {
	// success
	err := EvarDoesNotExist("NAME")
	assert.EqualError(t, err, `environment variable "NAME" does not exist`)
}

func TestFileAlreadyExists(t *testing.T) {
	// success
	err := FileAlreadyExists("/dire/name.extn")
	assert.EqualError(t, err, `file "name.extn" already exists`)
}

func TestFileDoesNotExist(t *testing.T) {
	// success
	err := FileDoesNotExist("/dire/name.extn")
	assert.EqualError(t, err, `file "name.extn" does not exist`)
}

func TestFileSystemError(t *testing.T) {
	// success
	err := FileSystemError("/dire/name.extn", errors.New("error"))
	assert.EqualError(t, err, `file "name.extn" had system error: error`)
}

func TestProgDoesNotExist(t *testing.T) {
	// success
	err := ProgDoesNotExist("name")
	assert.EqualError(t, err, `program "name" does not exist`)
}

func TestProgSystemError(t *testing.T) {
	// success
	err := ProgSystemError("name", errors.New("error"))
	assert.EqualError(t, err, `program "name" had system error: error`)
}
