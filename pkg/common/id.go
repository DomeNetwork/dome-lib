package common

import uuid "github.com/satori/go.uuid"

// ID for identifying items in a database which is currently
// UUIDv4 strings.
type ID string

// NewID returns a newly generated UUIDv4 identifier.
func NewID() ID {
	return ID(uuid.NewV4().String())
}
