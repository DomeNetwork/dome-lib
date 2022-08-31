package common

import "time"

// Now returns the current UTC time.
func Now() (t time.Time) {
	t = time.Now().UTC()
	return
}

// Unix returns the current UTC Unix timestamp seconds.
// Note: This does not include milliseconds.
func Unix() (t int64) {
	t = Now().Unix()
	return
}
