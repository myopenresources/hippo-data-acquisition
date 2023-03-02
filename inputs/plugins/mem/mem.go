package mem

import (
	"hippo-data-acquisition/commons/queue"
	"hippo-data-acquisition/config"
	"hippo-data-acquisition/inputs/input_collection"
)

type FsNotify struct {
}

func (f *FsNotify) InitPlugin(config config.InputConfig) {

}

func (f *FsNotify) PrepareCron() {

}

func (f *FsNotify) BeforeExeDataAcquisition() {

}

func (f *FsNotify) ExeDataAcquisition(dataQueue queue.Queue) {
	fields := make(map[string]interface{})
	tags := make(map[string]string)

	memInfo := GetMemPercent()
	if memInfo != nil {
		fields["memInfo"] = memInfo
	}

	dataQueue.PushData(fields, tags)

}
func (f *FsNotify) AfterExeDataAcquisition() {

}

func init() {
	input_collection.Add("mem", &FsNotify{})
}
