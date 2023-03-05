package desensitization

import (
	"encoding/json"
	"hippo-data-acquisition/commons/logger"
	"hippo-data-acquisition/commons/queue"
	"hippo-data-acquisition/config"
	"hippo-data-acquisition/processors/processors_collection"
	"reflect"
	"strings"
)

type Desensitization struct {
	keywords []map[string]interface{}
}

// InitPlugin 初始化参数
func (d *Desensitization) InitPlugin(config config.ProcessorConfig, pluginName string) {
	if len(d.keywords) <= 0 {
		var keywords []map[string]interface{}
		value := reflect.ValueOf(config.Params["keywords"])
		if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
			logger.LogError("desensitization", "脱敏处理器参数不是数组或切片！")
		}

		for i := 0; i < value.Len(); i++ {
			keywords = append(keywords, value.Index(i).Interface().(map[string]interface{}))
		}

		d.keywords = keywords
	}

}

// BeforeExeProcess  执行处理前
func (d *Desensitization) BeforeExeProcess() {

}

// ExeProcess  执行处理
func (d *Desensitization) ExeProcess(dataQueue queue.Queue) {
	for i := range dataQueue.GetDataList() {
		dataInfo := dataQueue.GetDataList()[i]
		dataByte, err := json.Marshal(&dataInfo)
		if err != nil {
			logger.LogError("desensitization", "脱敏处理器转换数据为json字符串失败："+err.Error())
			continue
		}
		jsonStr := string(dataByte)

		for i2 := range d.keywords {
			item := d.keywords[i2]
			jsonStr = strings.ReplaceAll(jsonStr, (item["value"]).(string), (item["desensitizationSymbol"]).(string))
		}

		var newDataInfo = queue.DataInfo{}
		err = json.Unmarshal([]byte(jsonStr), &newDataInfo)
		if err != nil {
			logger.LogError("desensitization", "脱敏处理器转换json字符为对象失败："+err.Error())
			continue
		}
		dataQueue.GetDataList()[i] = newDataInfo
	}
}

// AfterExeProcess  执行处理后
func (d *Desensitization) AfterExeProcess() {

}

func init() {
	processors_collection.Add("desensitization", &Desensitization{})
}
