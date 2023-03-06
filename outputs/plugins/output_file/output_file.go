package output_file

import (
	"encoding/json"
	"fmt"
	"hippo-data-acquisition/commons/logger"
	"hippo-data-acquisition/commons/queue"
	"hippo-data-acquisition/commons/utils"
	"hippo-data-acquisition/config"
	"hippo-data-acquisition/outputs/output_collection"
)

type OutputFile struct {
	filePath string
}

// InitPlugin 初始化参数
func (f *OutputFile) InitPlugin(config config.OutputConfig) {
	filePath, ok := config.Params["filePath"]
	if ok {
		f.filePath = utils.NewDateFilePath(filePath.(string), "")
	} else {
		logger.LogInfo("outputFile", "文件输出插件缺少参数：filePath")
	}

}

// BeforeExeOutput  执行输出前
func (f *OutputFile) BeforeExeOutput() {

}

// ExeOutput  执行输出
func (f *OutputFile) ExeOutput(dataInfo queue.DataInfo) {
	strByte, err := json.Marshal(&dataInfo)
	if err != nil {
		logger.LogInfo("outputFile", "输出数据转换成json字符串失败！")
	}

	utils.WriteStrToFile(f.filePath, utils.BytesToStr(strByte), "outputFile", func(log string) {
		fmt.Println(log)
	})

}

func init() {
	output_collection.Add("outputFile", &OutputFile{})
}
