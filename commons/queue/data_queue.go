package queue

import "errors"

type Queue interface {
	PopData() (error, DataInfo)
	PushData(dataBody map[string]interface{}, tag map[string]string)
	GetDataList() []DataInfo
}
type DataInfo struct {
	DataBody map[string]interface{}
	Tag      map[string]string
}

type DataQueue struct {
	dataList []DataInfo
}

func NewDataQueue() DataQueue {
	dataQueue := DataQueue{
		dataList: make([]DataInfo, 0),
	}
	return dataQueue
}

func (q *DataQueue) PushData(dataBody map[string]interface{}, tag map[string]string) {
	q.dataList = append(q.dataList, DataInfo{
		DataBody: dataBody,
		Tag:      tag,
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
