package fbhttp

import (
	"net/http"
	"time"

	"github.com/spf13/viper"
)

// FictionBase struct
type FictionBase struct {
	Message HTTP `json:"message"`
}

// HTTP struct
type HTTP struct {
	TypeKey      string    `json:"type_key"`
	StorageKey   string    `json:"storage_key"`
	TimeKey      time.Time `json:"time_key"`
	MonitorHTTP  string    `json:"monitor_http"`
	Status       float64   `json:"status"`
	ResponseTime float64   `json:"response_time"`
}

// GetResponseAndTime Get GetResponseData And GetResponseTime
func GetResponseAndTime() (*http.Response, float64, error) {
	start := time.Now()
	resp, err := http.Get(viper.GetString("externalMonitoring.http"))
	if err != nil {
		return nil, 0, err
	}
	elapsed := time.Since(start).Seconds()
	return resp, elapsed, nil
}

// InitKey set FictionBase Keys
func (fb FictionBase) InitKey() {
	fb.Message.TypeKey = "fbhttp.HTTP"
	fb.Message.StorageKey = "cloudwatch"
	fb.Message.TimeKey = time.Now()
	fb.Message.MonitorHTTP = viper.GetString("externalMonitoring.http")
}
