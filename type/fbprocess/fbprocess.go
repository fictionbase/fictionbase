package fbprocess

import (
	"fmt"
	"os"
	"time"

	ps "github.com/mitchellh/go-ps"
)

// FictionBase struct
type FictionBase struct {
	Message Processes `json:"message"`
}

// Processes struct
type Processes struct {
	TypeKey    string    `json:"type_key"`
	StorageKey string    `json:"storage_key"`
	TimeKey    time.Time `json:"time_key"`
	process    string
}

// Run GetResource And Send SQS
func (fb *FictionBase) Run() {
	pss, err := ps.Processes()
	for i, j := range pss {
		fmt.Println(i)
		fmt.Println(j)
	}
	fmt.Println(err)
	os.Exit(0)
}
