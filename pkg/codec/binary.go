package codec

import (
	"encoding/gob"
	"io"
)

var _ Codec = &Binary{}

// Binary provides GOB encoding operations from the
// standard Go encoding/gob library.
type Binary struct{}

// NewBinary returns a new Binary codec.
func NewBinary() *Binary {
	return &Binary{}
}

// Decode the provided interface object from the given reader or throw an error.
func (o *Binary) Decode(r io.Reader, v interface{}) (err error) {
	err = gob.NewDecoder(r).Decode(v)
	return
}

// Encode the provided interface into the writer or throw an error.
func (o *Binary) Encode(w io.Writer, v interface{}) (err error) {
	err = gob.NewEncoder(w).Encode(v)
	return
}
