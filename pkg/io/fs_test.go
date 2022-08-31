package io

import (
	"bytes"
	"testing"
)

var (
	fs     = NewFS()
	fsPath = "/tmp/dome-test.txt"
)

func TestFSWrite(t *testing.T) {
	r := bytes.NewBufferString("DOME")
	if err := fs.Write(fsPath, r); err != nil {
		t.Error(err)
		return
	}
}

func TestFSCheck(t *testing.T) {
	if !fs.Check(fsPath) {
		t.Errorf("unable to find test file: %s", fsPath)
		return
	}
}

func TestFSRead(t *testing.T) {
	w := new(bytes.Buffer)
	if err := fs.Read(fsPath, w); err != nil {
		t.Error(err)
		return
	}

	if w.String() != "DOME" {
		t.Errorf("file contents mismatch: %s != %s", "DOME", w.String())
	}
}
