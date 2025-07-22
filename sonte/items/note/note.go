// Package note implements the Note type and methods.
package note

import (
	"os"

	"github.com/gesedels/sonte/sonte/tools/file"
	"github.com/gesedels/sonte/sonte/tools/neat"
	"github.com/gesedels/sonte/sonte/tools/path"
)

// Note is a single plaintext note file in a directory.
type Note struct {
	Orig string
	Mode os.FileMode
}

// New returns a new Note.
func New(orig string, mode os.FileMode) *Note {
	return &Note{orig, mode}
}

// Delete "deletes" the Note by changing its extension to ".trash".
func (n *Note) Delete() error {
	return file.Delete(n.Orig)
}

// Exists returns true if the Note exists.
func (n *Note) Exists() bool {
	return file.Exists(n.Orig)
}

// Match returns true if the Note's name contains a case-insensitive substring.
func (n *Note) Match(text string) bool {
	return path.Match(n.Orig, text)
}

// Name returns the Note's file name.
func (n *Note) Name() string {
	name := path.Name(n.Orig)
	return neat.Name(name)
}

// Read returns the Note's body as a string.
func (n *Note) Read() (string, error) {
	body, err := file.Read(n.Orig)
	return neat.Body(body), err
}

// Rename changes the Note's name.
func (n *Note) Rename(name string) error {
	name = neat.Name(name)
	return file.Rename(n.Orig, name)
}

// Search returns true if the Note's body contains a case-insensitive substring.
func (n *Note) Search(text string) (bool, error) {
	return file.Search(n.Orig, text)
}

// Update overwrites the Note's body with a string.
func (n *Note) Update(body string) error {
	body = neat.Body(body)
	return file.Update(n.Orig, body, n.Mode)
}
