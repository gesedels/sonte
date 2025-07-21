package neat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
	// success
	body := Body("\nBody.\n")
	assert.Equal(t, "Body.\n", body)
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
