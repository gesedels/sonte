// Package file implements file system I/O functions.
package file

import (
	"errors"
	"os"
	"strings"

	"github.com/gesedels/sonte/sonte/tools/errs"
	"github.com/gesedels/sonte/sonte/tools/path"
)

// Create creates a new file containing a string.
func Create(dest, body string, mode os.FileMode) error {
	if Exists(dest) {
		return errs.FileAlreadyExists(dest)
	}

	if err := os.WriteFile(dest, []byte(body), mode); err != nil {
		return errs.FileSystemError(dest, err)
	}

	return nil
}

// Delete "deletes" an existing file by changing its extension to ".trash".
func Delete(orig string) error {
	if Exists(orig) {
		dire := path.Dire(orig)
		name := path.Name(orig)
		dest := path.Join(dire, name, ".trash")

		if err := os.Rename(orig, dest); err != nil {
			return errs.FileSystemError(orig, err)
		}
	}

	return nil
}

// Exists returns true if a file or directory exists.
func Exists(orig string) bool {
	_, err := os.Stat(orig)
	return !errors.Is(err, os.ErrNotExist)
}

// Read returns a file's body as a string.
func Read(orig string) (string, error) {
	if !Exists(orig) {
		return "", errs.FileDoesNotExist(orig)
	}

	bytes, err := os.ReadFile(orig)
	if err != nil {
		return "", errs.FileSystemError(orig, err)
	}

	return string(bytes), nil
}

// Rename moves a file to a different name.
func Rename(orig, name string) error {
	if !Exists(orig) {
		return errs.FileDoesNotExist(orig)
	}

	dire := path.Dire(orig)
	extn := path.Extn(orig)
	dest := path.Join(dire, name, extn)

	if err := os.Rename(orig, dest); err != nil {
		return errs.FileSystemError(orig, err)
	}

	return nil
}

// Search returns true if a file's body contains a substring.
func Search(orig, text string) (bool, error) {
	if !Exists(orig) {
		return false, errs.FileDoesNotExist(orig)
	}

	bytes, err := os.ReadFile(orig)
	if err != nil {
		return false, errs.FileSystemError(orig, err)
	}

	text = strings.ToLower(text)
	body := strings.ToLower(string(bytes))
	return strings.Contains(body, text), nil
}

// Update overwrites an existing file's body with a string.
func Update(orig, body string, mode os.FileMode) error {
	if !Exists(orig) {
		return errs.FileDoesNotExist(orig)
	}

	if err := os.WriteFile(orig, []byte(body), mode); err != nil {
		return errs.FileSystemError(orig, err)
	}

	return nil
}
