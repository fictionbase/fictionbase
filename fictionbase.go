package fictionbase

import (
	"path/filepath"

	"github.com/spf13/viper"
)

// SetViperConfig Set And Read ViperConfig
func SetViperConfig() {
	path := filepath.Join(
		"$GOPATH",
		"src",
		"github.com",
		"fictionbase",
		"fictionbase")
	viper.AddConfigPath(path)
	viper.SetConfigName("fictionbase")
	viper.ReadInConfig()
}
