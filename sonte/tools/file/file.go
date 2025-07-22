// Package file implements file system I/O functions.
package file

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gesedels/sonte/sonte/tools/path"
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

// Delete "deletes" an existing file by changing its extension to ".trash".
func Delete(orig string) error {
	if Exists(orig) {
		dire := path.Dire(orig)
		name := path.Name(orig)
		dest := path.Join(dire, name, ".trash")

		if err := os.Rename(orig, dest); err != nil {
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

// Rename changes a file's path to a different name.
func Rename(orig, name string) error {
	if !Exists(orig) {
		base := path.Base(orig)
		return fmt.Errorf("cannot rename file %q - does not exist", base)
	}

	dire := path.Dire(orig)
	extn := path.Extn(orig)
	dest := path.Join(dire, name, extn)

	if err := os.Rename(orig, dest); err != nil {
		base := path.Base(orig)
		return fmt.Errorf("cannot rename file %q - %w", base, err)
	}

	return nil
}

// Search returns true if a file's body contains a case-insensitive substring.
func Search(orig, text string) (bool, error) {
	if !Exists(orig) {
		base := path.Base(orig)
		return false, fmt.Errorf("cannot search file %q - does not exist", base)
	}

	bytes, err := os.ReadFile(orig)
	if err != nil {
		base := path.Base(orig)
		return false, fmt.Errorf("cannot search file %q - %w", base, err)
	}

	text = strings.ToLower(text)
	body := strings.ToLower(string(bytes))
	return strings.Contains(body, text), nil
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
