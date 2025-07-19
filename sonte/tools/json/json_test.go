package json

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	// setup
	var jval float64
	bytes := []byte(`123`)

	// success
	err := Decode(bytes, &jval)
	assert.Equal(t, 123.0, jval)
	assert.NoError(t, err)

	// error - cannot decode
	err = Decode(nil, nil)
	assert.EqualError(t, err, "cannot decode JSON byteslice")
}

func TestDecodeMap(t *testing.T) {
	// setup
	bytes := []byte(`{"a":"b"}`)

	// success
	jmap, err := DecodeMap(bytes)
	assert.Equal(t, map[string]any{"a": "b"}, jmap)
	assert.NoError(t, err)

	// error - cannot decode
	jmap, err = DecodeMap(nil)
	assert.Nil(t, jmap)
	assert.EqualError(t, err, "cannot decode JSON byteslice into map")
}

func TestEncode(t *testing.T) {
	// success
	bytes, err := Encode(123)
	assert.Equal(t, []byte(`123`), bytes)
	assert.NoError(t, err)

	// error - cannot encode
	bytes, err = Encode(make(chan int))
	assert.Nil(t, bytes)
	assert.EqualError(t, err, "cannot encode JSON value")
}
