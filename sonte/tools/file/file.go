// Package file implements file system I/O functions.
package file

import (
	"errors"
	"fmt"
	"os"
	"path"
)

// Create creates a new file containing a string.
func Create(dest, body string, mode os.FileMode) error {
	if Exists(dest) {
		base := path.Base(dest)
		return fmt.Errorf("cannot create file %q - already exists", base)
	}

	if err := os.WriteFile(dest, []byte(body), mode); err != nil {
		base := path.Base(dest)
		return fmt.Errorf("cannot create file %q - %w", base, err)
	}

	return nil
}

// Delete deletes an existing file.
func Delete(orig string) error {
	if Exists(orig) {
		if err := os.Remove(orig); err != nil {
			base := path.Base(orig)
			return fmt.Errorf("cannot delete file %q - %w", base, err)
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
		base := path.Base(orig)
		return "", fmt.Errorf("cannot read file %q - does not exist", base)
	}

	bytes, err := os.ReadFile(orig)
	if err != nil {
		base := path.Base(orig)
		return "", fmt.Errorf("cannot read file %q - %w", base, err)
	}

	return string(bytes), nil
}

// Update overwrites an existing file with a string.
func Update(orig, body string, mode os.FileMode) error {
	if !Exists(orig) {
		base := path.Base(orig)
		return fmt.Errorf("cannot update file %q - does not exist", base)
	}

	if err := os.WriteFile(orig, []byte(body), mode); err != nil {
		base := path.Base(orig)
		return fmt.Errorf("cannot update file %q - %w", base, err)
	}

	return nil
}
