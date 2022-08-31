package kv

import "time"

// KV sets an interface for getting and setting of
// values and keys.
type KV interface {
	Close() error
	Delete(string) error
	Get(string) (interface{}, error)
	Open() error
	Set(string, interface{}, time.Duration) error
}
