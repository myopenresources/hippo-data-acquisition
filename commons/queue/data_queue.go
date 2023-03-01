package queue

import (
	"encoding/json"
	"errors"
	"hippo-data-acquisition/commons/utils"
	"time"
)

type Queue interface {
	PopData() (error, DataInfo, string)
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
	if tag == nil {
		tag = make(map[string]string)
	}
	utils.Concat(tag, q.globalTag)

	tag["dataAcquisitionTime"] = time.Now().Format("2006-01-02 15:04:05")

	q.dataList = append(q.dataList, DataInfo{
		DataBody: dataBody,
		Tag:      tag,
	})
}

func (q *DataQueue) PopData() (error, DataInfo, string) {
	if len(q.dataList) == 0 {
		return errors.New("队列为空！"), DataInfo{}, ""
	}
	v := q.dataList[0]
	q.dataList = q.dataList[1:]

	strByte, err := json.Marshal(&v)
	if err != nil {
		return nil, v, ""
	}

	return nil, v, string(strByte)
}

func (q *DataQueue) GetDataList() []DataInfo {
	return q.dataList
}

func (q *DataQueue) SetDataList(list []DataInfo) {
	q.dataList = list
}
