package fbprocess

import (
	"time"

	"github.com/fictionbase/fictionbase"
	ps "github.com/mitchellh/go-ps"
)

// FictionBase struct
type FictionBase struct {
	Message Processes `json:"message"`
}

// Processes struct
type Processes struct {
	fictionbase.MessageBase
	Process string `json:"process"`
	Exists  bool   `json:"exists"`
}

// Run GetResource And Send SQS
func (fb *FictionBase) Run() {
	fb.Message.TypeKey = "fbprocess"
	fb.Message.StorageKey = "cloudwatch"
	for {
		time.Sleep(1 * time.Second)
		pss, err := ps.Processes()
		if err != nil {
			fictionbase.Logger.Error(err.Error())
			continue
		}
		// @TODO from config multiprocess
		fb.Message.Process = "httpd"
		fb.Message.Exists = true
		for _, process := range pss {
			if process.Executable() == fb.Message.Process {
				fb.Message.Exists = true
			}
		}
		// Set Time
		fb.Message.TimeKey = time.Now()
		err = fictionbase.SendFictionbaseMessage(fb)
		if err != nil {
			fictionbase.Logger.Error(err.Error())
		}
	}
}
