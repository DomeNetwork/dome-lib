package kv

import (
	"time"

	"github.com/domenetwork/dome-lib/pkg/log"
)

var _ KV = &Memory{}

// Memory interfaces with the famous KV storage engine Memory.
type Memory struct {
	items map[string]interface{}
}

// Close the connection.
func (kv *Memory) Close() (err error) {
	log.D("memory", "close")
	return
}

// Delete the item from the store.
func (kv *Memory) Delete(key string) (err error) {
	log.D("memory", "delete", key)
	delete(kv.items, key)
	return
}

// Get the value that matches the provided key from Memory.
func (kv *Memory) Get(key string) (value interface{}, err error) {
	log.D("memory", "get", key)
	value = kv.items[key]
	return
}

// Open the connection to Memory.
func (kv *Memory) Open() (err error) {
	log.D("memory", "open")
	kv.items = map[string]interface{}{}
	return
}

// Set the provided value to the provided key in Memory.
func (kv *Memory) Set(key string, value interface{}, ttl time.Duration) (err error) {
	log.D("memory", "set", key, value)
	kv.items[key] = value
	return
}
