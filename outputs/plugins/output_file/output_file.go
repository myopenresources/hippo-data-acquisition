package output_file

import (
	"encoding/json"
	"hippo-data-acquisition/commons/logger"
	"hippo-data-acquisition/commons/queue"
	"hippo-data-acquisition/config"
	"hippo-data-acquisition/outputs/output_collection"
	"io"
	"os"
)

type OutputFile struct {
	filePath string
}

// InitPlugin 初始化参数
func (f *OutputFile) InitPlugin(config config.OutputConfig) {
	filePath, ok := config.Params["filePath"]
	if ok {
		f.filePath = filePath.(string)
	} else {
		logger.LogInfo("outputFile", "文件输入插件缺少参数：filePath")
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
	writeDataToFile(f.filePath, string(strByte))

}

func writeDataToFile(filePath string, dataJson string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0644)
	if err != nil {
		logger.LogInfo("outputFile", "创建输入文件对象失败："+err.Error())
	} else {
		n, _ := file.Seek(0, io.SeekEnd)
		_, err = file.WriteAt([]byte(dataJson+"\n"), n)
		logger.LogInfo("outputFile", "输出数据："+dataJson+"到"+filePath)
	}
	defer file.Close()
}

func init() {
	output_collection.Add("outputFile", &OutputFile{})
}
