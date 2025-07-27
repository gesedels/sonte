package file

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/gesedels/sonte/sonte/tools/test"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	// setup
	dire := test.MockDire(t)
	dest := filepath.Join(dire, "name.extn")

	// success
	err := Create(dest, "Body.\n", 0666)
	test.AssertFile(t, dest, "Body.\n")
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")
	dest := strings.Replace(orig, ".extn", ".trash", 1)

	// success
	err := Delete(orig)
	assert.NoFileExists(t, orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)
}

func TestExists(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")
	nope := "/dire/nope.txt"

	// success - true
	ok := Exists(orig)
	assert.True(t, ok)

	// success - false
	ok = Exists(nope)
	assert.False(t, ok)
}

func TestRead(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success
	body, err := Read(orig)
	assert.Equal(t, test.MockData["alpha.extn"], body)
	assert.NoError(t, err)
}

func TestRename(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")
	dest := strings.Replace(orig, "alpha", "name", 1)

	// success
	err := Rename(orig, "name")
	assert.NoFileExists(t, orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)

}

func TestSearch(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success - true
	ok, err := Search(orig, "ALPH")
	assert.True(t, ok)
	assert.NoError(t, err)

	// success - false
	ok, err = Search(orig, "nope")
	assert.False(t, ok)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success
	err := Update(orig, "Body.\n", 0666)
	test.AssertFile(t, orig, "Body.\n")
	assert.NoError(t, err)
}
