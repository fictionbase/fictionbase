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
	process string
	exists  bool
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
		fb.Message.process = "httpd"
		fb.Message.exists = false
		for _, process := range pss {
			if process.Executable() == fb.Message.process {
				fb.Message.exists = true
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
