// Package book implements the Book type and methods.
package book

import (
	"fmt"
	"os"

	"github.com/gesedels/sonte/sonte/items/note"
	"github.com/gesedels/sonte/sonte/tools/clui"
	"github.com/gesedels/sonte/sonte/tools/file"
	"github.com/gesedels/sonte/sonte/tools/neat"
	"github.com/gesedels/sonte/sonte/tools/path"
)

// Book is a single directory containing multiple Notes.
type Book struct {
	Dire string
	Extn string
	Mode os.FileMode
}

// New returns a new Book.
func New(dire, extn string, mode os.FileMode) *Book {
	dire = neat.Dire(dire)
	extn = neat.Extn(extn)
	return &Book{dire, extn, mode}
}

// NewEnv returns a new Book using environment variables.
func NewEnv(denv, eenv string, mode os.FileMode) (*Book, error) {
	dire, err := clui.Env(denv)
	if err != nil {
		return nil, err
	}

	extn, err := clui.Env(eenv)
	if err != nil {
		return nil, err
	}

	return New(dire, extn, mode), nil
}

// Create creates and returns a new Note in the Book.
func (b *Book) Create(name, body string) (*note.Note, error) {
	name = neat.Name(name)
	body = neat.Body(body)
	dest := path.Join(b.Dire, name, b.Extn)
	if err := file.Create(dest, body, b.Mode); err != nil {
		return nil, err
	}

	return note.New(dest, b.Mode), nil
}

// Exists returns true if the Book's directory exists.
func (b *Book) Exists() bool {
	return file.Exists(b.Dire)
}

// Filter returns all existing Notes in the Book that succeed a filter function.
func (b *Book) Filter(ffun func(*note.Note) (bool, error)) ([]*note.Note, error) {
	var notes []*note.Note

	for _, note := range b.List() {
		ok, err := ffun(note)
		switch {
		case ok:
			notes = append(notes, note)
		case err != nil:
			return nil, err
		}
	}

	return notes, nil
}

// Get returns an existing Note from the Book.
func (b *Book) Get(name string) (*note.Note, error) {
	name = neat.Name(name)
	orig := path.Join(b.Dire, name, b.Extn)

	if !file.Exists(orig) {
		base := path.Base(orig)
		return nil, fmt.Errorf("cannot locate file %q - does not exist", base)
	}

	return note.New(orig, b.Mode), nil
}

// GetOrCreate returns a newly created or existing Note from the Book.
func (b *Book) GetOrCreate(name, body string) (*note.Note, error) {
	name = neat.Name(name)
	body = neat.Body(body)
	orig := path.Join(b.Dire, name, b.Extn)

	if !file.Exists(orig) {
		if err := file.Create(orig, body, b.Mode); err != nil {
			return nil, err
		}
	}

	return note.New(orig, b.Mode), nil
}

// List returns all existing Notes in the Book in alphabetical order.
func (b *Book) List() []*note.Note {
	var notes []*note.Note

	for _, orig := range path.Glob(b.Dire, b.Extn) {
		notes = append(notes, note.New(orig, b.Mode))
	}

	return notes
}

// Match returns all existing Notes in the Book with names containing a substring.
func (b *Book) Match(text string) []*note.Note {
	notes, _ := b.Filter(func(note *note.Note) (bool, error) {
		return note.Match(text), nil
	})

	return notes
}

// Search returns all existing Notes in the Book with bodies containing a substring.
func (b *Book) Search(text string) ([]*note.Note, error) {
	return b.Filter(func(note *note.Note) (bool, error) {
		return note.Search(text)
	})
}
