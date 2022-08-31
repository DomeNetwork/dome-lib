package codec

import (
	"encoding/json"
	"io"
)

var _ Codec = &JSON{}

// JSON provides JSON encoding operations from the
// standard Go encoding/json library.
type JSON struct{}

// NewJSON returns a new instance of the JSON codec.
func NewJSON() *JSON {
	return &JSON{}
}

// Decode the provided interface object from the given reader or throw an error.
func (o *JSON) Decode(r io.Reader, v interface{}) (err error) {
	err = json.NewDecoder(r).Decode(v)
	return
}

// Encode the provided interface into the writer or throw an error.
func (o *JSON) Encode(w io.Writer, v interface{}) (err error) {
	err = json.NewEncoder(w).Encode(v)
	return
}
