package path

import (
	"path/filepath"
	"testing"

	"github.com/gesedels/sonte/sonte/tools/test"
	"github.com/stretchr/testify/assert"
)

func TestBase(t *testing.T) {
	// success
	base := Base("/dire/name.extn")
	assert.Equal(t, "name.extn", base)
}

func TestDire(t *testing.T) {
	// success
	dire := Dire("/dire/name.extn")
	assert.Equal(t, "/dire", dire)
}

func TestExtn(t *testing.T) {
	// success - with extension
	extn := Extn("/dire/name.extn")
	assert.Equal(t, ".extn", extn)

	// success - with empty extension
	extn = Extn("/dire/name.")
	assert.Equal(t, ".", extn)

	// success - without extension
	extn = Extn("/dire/name")
	assert.Equal(t, "", extn)
}

func TestGlob(t *testing.T) {
	// setup
	dire := test.MockDire(t)

	// success
	origs := Glob(dire, ".extn")
	assert.Equal(t, []string{
		filepath.Join(dire, "alpha.extn"),
		filepath.Join(dire, "bravo.extn"),
	}, origs)
}

func TestJoin(t *testing.T) {
	// success
	dest := Join("/dire", "name", ".extn")
	assert.Equal(t, "/dire/name.extn", dest)
}

func TestMatch(t *testing.T) {
	// success - true
	ok := Match("/dire/name.extn", "nam")
	assert.True(t, ok)

	// success - false
	ok = Match("/dire/name.extn", "nope")
	assert.False(t, ok)
}

func TestName(t *testing.T) {
	// success - with extension
	name := Name("/dire/name.extn")
	assert.Equal(t, "name", name)

	// success - with empty extension
	name = Name("/dire/name.")
	assert.Equal(t, "name", name)

	// success - without extension
	name = Name("/dire/name")
	assert.Equal(t, "name", name)
}
