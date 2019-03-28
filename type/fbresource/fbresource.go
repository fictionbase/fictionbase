package fbresource

import (
	"fmt"
	"time"

	"github.com/fictionbase/fictionbase"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

// FictionBase struct
type FictionBase struct {
	Message Resources `json:"message"`
}

// Resources struct
type Resources struct {
	fictionbase.MessageBase
	Memory  *mem.VirtualMemoryStat `json:"memory"`
	CPU     []cpu.InfoStat         `json:"cpu"`
	LoadAvg *load.AvgStat          `json:"load_avg"`
	Host    *host.InfoStat         `json:"host"`
	Disk    *disk.UsageStat        `json:"disk"`
}

// Run GetResource And Send SQS
func (fb *FictionBase) Run() {
	var err error
	fb.Message.TypeKey = "fbresource"
	fb.Message.StorageKey = "cloudwatch"
	for {
		time.Sleep(1 * time.Second)
		// Memory
		fb.Message.Memory, err = mem.VirtualMemory()
		if err != nil {
			fmt.Println(err)
			fb.Message.Memory = nil
		}
		// CPU
		fb.Message.CPU, err = cpu.Info()
		if err != nil {
			fb.Message.CPU = nil
		}
		// LoadAvg
		fb.Message.LoadAvg, err = load.Avg()
		if err != nil {
			fmt.Println(err)
			fb.Message.LoadAvg = nil
		}
		// Host
		fb.Message.Host, err = host.Info()
		if err != nil {
			fmt.Println(err)
			fb.Message.Host = nil
		}
		// Disk
		fb.Message.Disk, err = disk.Usage("/")
		if err != nil {
			fmt.Println(err)
			fb.Message.Disk = nil
		}
		// Set TimeC
		fb.Message.TimeKey = time.Now()
		err = fictionbase.SendFictionbaseMessage(fb)
		if err != nil {
			fmt.Println(err)
		}
	}
}
