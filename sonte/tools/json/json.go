// Package json implements JSON encoding and decoding functions.
package json

import (
	"encoding/json"

	"github.com/gesedels/sonte/sonte/tools/errs"
)

// Decode decodes a JSON byteslice into a pointer value.
func Decode(bytes []byte, jval any) error {
	if err := json.Unmarshal(bytes, jval); err != nil {
		return errs.Wrap(err, "cannot decode JSON byteslice")
	}

	return nil
}

// DecodeMap decodes a JSON byteslice into a string-any map.
func DecodeMap(bytes []byte) (map[string]any, error) {
	var jmap = make(map[string]any)
	if err := json.Unmarshal(bytes, &jmap); err != nil {
		return nil, errs.Wrap(err, "cannot decode JSON byteslice into map")
	}

	return jmap, nil
}

// Encode returns a JSON byteslice from a value.
func Encode(jval any) ([]byte, error) {
	bytes, err := json.MarshalIndent(jval, "", "  ")
	if err != nil {
		return nil, errs.Wrap(err, "cannot encode JSON value")
	}

	return bytes, nil
}
