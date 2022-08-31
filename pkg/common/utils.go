package common

import (
	"encoding/json"
)

// MapToInterface is a convenience function that will take in a map and convert it to the provided
// interface.  This is commonly used to convert a JSON map an internal type.
func MapToInterface(m map[string]interface{}, v interface{}) (err error) {
	var data []byte
	if data, err = json.Marshal(m); err != nil {
		return
	}

	err = json.Unmarshal(data, v)
	return
}
