package out_file

import (
	"encoding/json"
	"fmt"
	"hippo-data-acquisition/commons/logger"
	"hippo-data-acquisition/commons/queue"
	"hippo-data-acquisition/commons/utils"
	"hippo-data-acquisition/config"
	"hippo-data-acquisition/outputs/output_collection"
)

type OutFile struct {
	filePath string
}

// InitPlugin 初始化参数
func (f *OutFile) InitPlugin(config config.OutputConfig) {
	filePath, ok := config.Params["filePath"]
	if ok {
		f.filePath = utils.NewDateFilePath(filePath.(string), "")
	} else {
		logger.LogInfo("outFile", "文件输出插件缺少参数：filePath")
	}

}

// BeforeExeOutput  执行输出前
func (f *OutFile) BeforeExeOutput() {

}

// ExeOutput  执行输出
func (f *OutFile) ExeOutput(dataInfo queue.DataInfo) {
	strByte, err := json.Marshal(&dataInfo)
	if err != nil {
		logger.LogInfo("outFile", "输出数据转换成json字符串失败！")
	}

	utils.WriteStrToFile(f.filePath, utils.BytesToStr(strByte), "OutFile", func(log string) {
		fmt.Println(log)
	})

}

func init() {
	output_collection.Add("outFile", &OutFile{})
}
