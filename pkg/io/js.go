//go:build wasm
// +build wasm

package io

import (
	"bytes"
	"fmt"
	"io"
	"syscall/js"

	"github.com/domenetwork/dome-lib/pkg/log"
)

var (
	_            IO = &JS{}
	localStorage js.Value
	ErrNotLoaded = fmt.Errorf("it seems like localStorage is not loaded")
)

// JS for interacting with the local storage in the browser.
type JS struct{}

// NewJS will return a JS instance.
func NewJS() (o *JS) {
	localStorage = js.Global().Get("localStorage")
	o = &JS{}
	return
}

// Check that the path exists.
func (o *JS) Check(key string) bool {
	log.D("io", "js", "check", key)
	if !localStorage.Truthy() {
		log.E("io", "js", "check", ErrNotLoaded)
		return false
	}

	v := localStorage.Get(key)
	return v.Truthy()
}

// Read the local storage at provided key and return its data or
// return an error.
func (o *JS) Read(key string, w io.Writer) (err error) {
	log.D("io", "js", "read", key)
	if !localStorage.Truthy() {
		err = ErrNotLoaded
		return
	}

	v := localStorage.Get(key)
	if !v.Truthy() {
		return
	}

	_, err = w.Write([]byte(v.String()))
	return
}

// Write the provided string to the given key or return error.
func (o *JS) Write(key string, r io.Reader) (err error) {
	log.D("io", "js", "write", key, localStorage)
	if !localStorage.Truthy() {
		err = ErrNotLoaded
		return
	}

	b := new(bytes.Buffer)
	if _, err = b.ReadFrom(r); err != nil {
		return
	}

	localStorage.Set(key, b.String())
	return
}
