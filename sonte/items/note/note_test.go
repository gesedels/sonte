package note

import (
	"os"
	"strings"
	"testing"

	"github.com/gesedels/sonte/sonte/tools/test"
	"github.com/stretchr/testify/assert"
)

func mockNote(t *testing.T) *Note {
	orig := test.MockFile(t, "alpha.extn")
	return New(orig, 0666)
}

func TestNew(t *testing.T) {
	// success
	note := mockNote(t)
	assert.Contains(t, note.Orig, "alpha.extn")
	assert.Equal(t, os.FileMode(0666), note.Mode)
}

func TestDelete(t *testing.T) {
	// setup
	note := mockNote(t)
	dest := strings.Replace(note.Orig, ".extn", ".trash", 1)

	// success
	err := note.Delete()
	assert.NoFileExists(t, note.Orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)
}

func TestExists(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	ok := note.Exists()
	assert.True(t, ok)
}

func TestMatch(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	ok := note.Match("ALPH")
	assert.True(t, ok)
}

func TestName(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	name := note.Name()
	assert.Equal(t, "alpha", name)
}

func TestRead(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	body, err := note.Read()
	assert.Equal(t, test.MockData["alpha.extn"], body)
	assert.NoError(t, err)
}

func TestRename(t *testing.T) {
	// setup
	note := mockNote(t)
	dest := strings.Replace(note.Orig, "alpha", "name", 1)

	// success
	err := note.Rename("name")
	assert.NoFileExists(t, note.Orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)
}

func TestSearch(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	ok, err := note.Search("ALPH")
	assert.True(t, ok)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	// setup
	note := mockNote(t)

	// success
	err := note.Update("Body.\n")
	test.AssertFile(t, note.Orig, "Body.\n")
	assert.NoError(t, err)
}
