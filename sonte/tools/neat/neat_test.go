package neat

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
	// success
	body := Body("\nBody.\n")
	assert.Equal(t, "Body.\n", body)
}

func TestDate(t *testing.T) {
	// setup
	want := time.Date(2001, time.February, 3, 0, 0, 0, 0, time.UTC)
	zero := time.Unix(0, 0)

	// success - valid date
	tobj := Date("2001-02-03")
	assert.Equal(t, want, tobj)

	// success - invalid date
	tobj = Date("")
	assert.Equal(t, zero, tobj)
}

func TestName(t *testing.T) {
	// success
	name := Name("\tNAME_123!!!\n")
	assert.Equal(t, "name-123", name)
}

func TestPath(t *testing.T) {
	// success
	path := Path("/././path")
	assert.Equal(t, "/path", path)
}
