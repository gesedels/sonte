package neat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
	// success
	body := Body("\tBody.\n")
	assert.Equal(t, "Body.\n", body)
}

func TestDire(t *testing.T) {
	// success
	dire := Dire("\t/././dire\n")
	assert.Equal(t, "/dire", dire)
}

func TextExtn(t *testing.T) {
	// success - with dot
	extn := Extn("\t.EXTN\n")
	assert.Equal(t, ".extn", extn)

	// success - without dot
	extn = Extn("\tEXTN\n")
	assert.Equal(t, ".extn", extn)
}

func TestName(t *testing.T) {
	// success
	name := Name("\tNAME_123!!!\n")
	assert.Equal(t, "name-123", name)
}
