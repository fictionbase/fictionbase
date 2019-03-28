package fbprocess

import (
	"fmt"
	"time"

	"github.com/fictionbase/agent"
	"github.com/fictionbase/fictionbase"
	ps "github.com/mitchellh/go-ps"
)

// FictionBase struct
type FictionBase struct {
	Message Processes `json:"message"`
}

// Processes struct
type Processes struct {
	agent.MessageBase
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
			fmt.Println(err)
			continue
		}
		// @TODO from config
		fb.Message.process = "httpd"
		fb.Message.exists = false
		for _, process := range pss {
			if process == fb.Message.process {
				fb.Message.exists = true
			}
		}
		err = fictionbase.SendFictionbaseMessage(fb)
		if err != nil {
			fmt.Println(err)
		}
	}
}
