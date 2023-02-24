package queue

import (
	"errors"
	"golang.org/x/exp/maps"
)

type Queue interface {
	PopData() (error, DataInfo)
	PushData(dataBody map[string]interface{}, tag map[string]string)
	GetDataList() []DataInfo
	SetDataList(list []DataInfo)
}
type DataInfo struct {
	DataBody map[string]interface{}
	Tag      map[string]string
}

type DataQueue struct {
	dataList  []DataInfo
	globalTag map[string]string
}

func NewDataQueue(globalTag map[string]string) DataQueue {
	initTag := globalTag
	if initTag == nil {
		initTag = make(map[string]string)
	}
	dataQueue := DataQueue{
		dataList:  make([]DataInfo, 0),
		globalTag: initTag,
	}
	return dataQueue
}

func (q *DataQueue) PushData(dataBody map[string]interface{}, tag map[string]string) {
	if tag != nil {
		maps.Copy(q.globalTag, tag)
	}

	q.dataList = append(q.dataList, DataInfo{
		DataBody: dataBody,
		Tag:      q.globalTag,
	})
}

func (q *DataQueue) PopData() (error, DataInfo) {
	if len(q.dataList) == 0 {
		return errors.New("队列为空！"), DataInfo{}
	}
	v := q.dataList[0]
	q.dataList = q.dataList[1:]
	return nil, v
}

func (q *DataQueue) GetDataList() []DataInfo {
	return q.dataList
}

func (q *DataQueue) SetDataList(list []DataInfo) {
	q.dataList = list
}
