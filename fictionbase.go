package fictionbase

import (
	"path/filepath"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// MessageBase fictionBase
type MessageBase struct {
	TypeKey    string    `json:"type_key"`
	StorageKey string    `json:"storage_key"`
	TimeKey    time.Time `json:"time_key"`
}

var (
	logger *zap.Logger
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

func init() {
	SetViperConfig()
	logger, _ = zap.NewProduction()
}
