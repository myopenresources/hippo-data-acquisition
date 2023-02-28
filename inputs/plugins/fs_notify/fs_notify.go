package fs_notify

import (
	"hippo-data-acquisition/commons/queue"
	"hippo-data-acquisition/config"
	"hippo-data-acquisition/inputs/input_collection"
	"time"
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

	fields["msg"] = "this is msg"
	tags := make(map[string]string)

	tags["test222"] = "250"
	tags["test"] = "sdffsafsdfsdfsdf"
	dataQueue.PushData(fields, tags)

	i := 0
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			fields["msg"] = "this is msg"
			tags := make(map[string]string)
			dataQueue.PushData(fields, tags)
			i++
			if i > 5 {
				return
			}
		}
	}

}
func (f *FsNotify) AfterExeDataAcquisition() {

}

func init() {
	input_collection.Add("fsNotify", &FsNotify{})
}
