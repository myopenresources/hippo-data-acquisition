package fs_notify

import (
	"hippo-data-acquisition/commons/queue"
	"hippo-data-acquisition/inputs/input_collection"
)

type FsNotify struct {
}

func (f *FsNotify) ExeDataAcquisition(dataQueue queue.Queue) {
	fields := make(map[string]interface{})
	fields["msg"] = "this is msg"

	tags := make(map[string]string)
	dataQueue.PushData(fields, tags)
}

func init() {
	input_collection.Add("fsNotify", &FsNotify{})
}
