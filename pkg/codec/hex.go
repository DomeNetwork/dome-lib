package codec

import (
	"bytes"
	"encoding/hex"
	"io"
	"strings"
)

var _ Codec = &Hex{}

// Hex provides hexadecimal encoding operations from the
// go-ethereum library.  Binary encoding takes place around encoding to
// hexadecimal.
type Hex struct{}

// NewHex returns a new instance of the Hex codec.
func NewHex() *Hex {
	return &Hex{}
}

// Decode the provided interface object from the given reader or throw an error.
func (o *Hex) Decode(r io.Reader, v interface{}) (err error) {
	sb := new(strings.Builder)
	if _, err = io.Copy(sb, r); err != nil {
		return
	}

	bin := new(Binary)
	r = bytes.NewBufferString(sb.String())
	err = bin.Decode(r, v)
	return
}

// Encode the provided interface into the writer or throw an error.
func (o *Hex) Encode(w io.Writer, v interface{}) (err error) {
	bin := new(Binary)
	r := new(bytes.Buffer)
	if err = bin.Decode(r, v); err != nil {
		return
	}

	s := make([]byte, 0)
	hex.Encode(s, r.Bytes())
	_, err = io.Copy(w, bytes.NewReader(s))
	return
}
