package cfg

import "fmt"

// ServiceURL returns the URL configured for the provided service name.
// It provides a dynamic reference to configuration key specific to
// API services.
func ServiceURL(name string) (url string) {
	key := fmt.Sprintf("sdk.%s", name)
	url = Str(key)
	return
}
