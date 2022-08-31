package codec

import "io"

// Codec provides an interface for all encoding options in Dome.
// Codecs are the encoding/decoding mechanism.
type Codec interface {
	Decode(io.Reader, interface{}) error
	Encode(io.Writer, interface{}) error
}
