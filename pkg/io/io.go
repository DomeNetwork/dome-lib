package io

import "io"

// IO provides an interface for writing data to a persistent
// storage location.
type IO interface {
	Check(string) bool
	Read(string, io.Writer) error
	Write(string, io.Reader) error
}
