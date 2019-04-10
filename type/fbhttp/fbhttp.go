package fbhttp

import (
	"time"

	"github.com/fictionbase/fictionbase"
	"github.com/spf13/viper"
)

// FictionBase struct
type FictionBase struct {
	Message HTTP `json:"message"`
}

// HTTP struct
type HTTP struct {
	fictionbase.MessageBase
	MonitorHTTP  string  `json:"monitor_http"`
	Status       float64 `json:"status"`
	ResponseTime float64 `json:"response_time"`
}

// InitKey set FictionBase Keys
func (fb FictionBase) InitKey() {
	fb.Message.TypeKey = "fbhttp.HTTP"
	fb.Message.StorageKey = "cloudwatch"
	fb.Message.TimeKey = time.Now()
	fb.Message.MonitorHTTP = viper.GetString("externalMonitoring.http")
}
