package cfg

import (
	"io"

	"github.com/spf13/viper"
)

// Bool value returned for the provided key.
func Bool(key string) (value bool) {
	value = viper.GetBool(key)
	return
}

// Int value returned from the provided key.
func Int(key string) (value int) {
	value = viper.GetInt(key)
	return
}

// Setup Viper with the provider reader and config type.  The reader
// should contain the contents in a format that matches the config type.
// The currently supported config types are JSON and YAML.
func Setup(r io.Reader, configType string) (err error) {
	viper.SetConfigType(configType)
	err = viper.ReadConfig(r)
	return
}

// Str is a string value returned for the provided key.
func Str(key string) (value string) {
	value = viper.GetString(key)
	return
}
