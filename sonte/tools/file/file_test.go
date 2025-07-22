package file

import (
	"path/filepath"
	"testing"

	"github.com/gesedels/sonte/sonte/tools/test"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	// setup
	dire := test.MockDire(t)
	dest := filepath.Join(dire, "create.extn")

	// success
	err := Create(dest, "Body.\n", 0666)
	test.AssertFile(t, dest, "Body.\n")
	assert.NoError(t, err)

	// error - already exists
	err = Create(dest, "Body.\n", 0666)
	assert.EqualError(t, err, `cannot create file "create.extn" - already exists`)
}

func TestDelete(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success
	err := Delete(orig)
	assert.NoFileExists(t, orig)
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
	nope := "/dire/nope.txt"

	// success
	body, err := Read(orig)
	assert.Equal(t, test.MockData["alpha.extn"], body)
	assert.NoError(t, err)

	// error - does not exist
	body, err = Read(nope)
	assert.Empty(t, body)
	assert.EqualError(t, err, `cannot read file "nope.txt" - does not exist`)
}

func TestUpdate(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")
	nope := "/dire/nope.txt"

	// success
	err := Update(orig, "Body.\n", 0666)
	test.AssertFile(t, orig, "Body.\n")
	assert.NoError(t, err)

	// error - does not exist
	err = Update(nope, "Body.\n", 0666)
	assert.EqualError(t, err, `cannot update file "nope.txt" - does not exist`)
}
