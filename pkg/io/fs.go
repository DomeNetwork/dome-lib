package io

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/domenetwork/dome-lib/pkg/log"
)

var _ IO = &FS{}

// FS allows writing to the file system.
type FS struct{}

// NewFS will return a FS instance for interacting with the local
// file system.
func NewFS() (o *FS) {
	o = &FS{}
	return
}

// Check that the path exists.
func (o *FS) Check(path string) bool {
	log.D("io", "fs", "check", path)
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Read the contents of a file at provided path and return data or error.
func (o *FS) Read(path string, w io.Writer) (err error) {
	log.D("io", "fs", "read", path)
	var b []byte
	if b, err = os.ReadFile(path); err != nil {
		return
	}

	_, err = w.Write(b)
	return
}

// Write the provided data to the given path or error.
func (o *FS) Write(path string, r io.Reader) (err error) {
	log.D("io", "fs", "write", path)
	var data []byte
	if data, err = ioutil.ReadAll(r); err != nil {
		return
	}

	err = ioutil.WriteFile(path, data, 0644)
	return
}
