package cpu

import (
	"hippo-data-acquisition/commons/queue"
	"hippo-data-acquisition/config"
	"hippo-data-acquisition/inputs/input_collection"
	"runtime"
)

type Cpu struct {
}

func (c *Cpu) InitPlugin(config config.InputConfig) {

}

func (c *Cpu) PrepareCron() {

}

func (c *Cpu) BeforeExeDataAcquisition() {

}

func (c *Cpu) ExeDataAcquisition(dataQueue queue.Queue) {
	fields := make(map[string]interface{})

	fields["cpuCore"] = runtime.NumCPU()
	tags := make(map[string]string)

	dataQueue.PushData(fields, tags)
}
func (c *Cpu) AfterExeDataAcquisition() {

}

func init() {
	input_collection.Add("cpu", &Cpu{})
}
