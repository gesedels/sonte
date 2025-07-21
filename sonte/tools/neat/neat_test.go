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

func TestHash(t *testing.T) {
	// success
	hash := Hash("test")
	assert.Equal(t, "n4bQgYhMfWWaL-qgxVrQFaO_TxsrC4Is0V1sFbDwCgg", hash)
}

func TestName(t *testing.T) {
	// success
	name := Name("\tNAME_123!!!\n")
	assert.Equal(t, "name-123", name)
}

func TestTime(t *testing.T) {
	// setup
	want := time.Unix(1234567890, 0)

	// success
	tobj := Time("\t1234567890\n")
	assert.Equal(t, want, tobj)
}

func TestUnix(t *testing.T) {
	// setup
	tobj := time.Unix(1234567890, 0)

	// success
	unix := Unix(tobj)
	assert.Equal(t, "1234567890", unix)
}
