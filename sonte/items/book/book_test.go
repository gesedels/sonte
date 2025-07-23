package book

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gesedels/sonte/sonte/items/note"
	"github.com/gesedels/sonte/sonte/tools/test"
	"github.com/stretchr/testify/assert"
)

func mockBook(t *testing.T) *Book {
	dire := test.MockDire(t)
	return New(dire, ".extn", 0666)
}

func TestNew(t *testing.T) {
	// success
	book := mockBook(t)
	assert.NotEmpty(t, book.Dire)
	assert.Equal(t, ".extn", book.Extn)
	assert.Equal(t, os.FileMode(0666), book.Mode)
}

func TestNewEnv(t *testing.T) {
	// setup
	os.Setenv("DIRE", mockBook(t).Dire)
	os.Setenv("EXTN", ".extn")

	// success
	book, err := NewEnv("DIRE", "EXTN", 0666)
	assert.NotEmpty(t, book.Dire)
	assert.Equal(t, ".extn", book.Extn)
	assert.Equal(t, os.FileMode(0666), book.Mode)
	assert.NoError(t, err)
}

func TestCreate(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	note, err := book.Create("name", "Body.\n")
	test.AssertFile(t, note.Orig, "Body.\n")
	assert.NoError(t, err)
}

func TestFilter(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	notes, err := book.Filter(func(note *note.Note) (bool, error) {
		return note.Name() == "alpha", nil
	})

	assert.Len(t, notes, 1)
	assert.Equal(t, "alpha", notes[0].Name())
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	book := mockBook(t)
	orig := filepath.Join(book.Dire, "alpha.extn")

	// success
	note, err := book.Get("alpha")
	assert.Equal(t, orig, note.Orig)
	assert.NoError(t, err)

	// error - does not exist
	note, err = book.Get("nope")
	assert.Nil(t, note)
	assert.EqualError(t, err, `cannot locate file "nope.extn" - does not exist`)
}

func TestGetOrCreate(t *testing.T) {
	// setup
	book := mockBook(t)

	// success - new note
	note, err := book.GetOrCreate("name", "Body.\n")
	test.AssertFile(t, note.Orig, "Body.\n")
	assert.NoError(t, err)

	// success - existing note
	note, err = book.GetOrCreate("name", "")
	test.AssertFile(t, note.Orig, "Body.\n")
	assert.NoError(t, err)
}

func TestList(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	notes := book.List()
	assert.Len(t, notes, 2)
	assert.Equal(t, "alpha", notes[0].Name())
	assert.Equal(t, "bravo", notes[1].Name())
}

func TestMatch(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	notes := book.Match("ALPH")
	assert.Len(t, notes, 1)
	assert.Equal(t, "alpha", notes[0].Name())
}

func TestSearch(t *testing.T) {
	// setup
	book := mockBook(t)

	// success
	notes, err := book.Search("ALPH")
	assert.Len(t, notes, 1)
	assert.Equal(t, "alpha", notes[0].Name())
	assert.NoError(t, err)
}
