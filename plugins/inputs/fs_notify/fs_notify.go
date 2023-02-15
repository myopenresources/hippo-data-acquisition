package fs_notify

import (
	"hippo-data-acquisition/commons/queue"
)

type FsNotify struct {
}

func (f *FsNotify) ExeDataAcquisition(dataQueue queue.Queue) {
	fields := make(map[string]interface{})
	fields["msg"] = "this is msg"

	tags := make(map[string]string)
	dataQueue.PushData(fields, tags)
}
