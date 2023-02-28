package output_file

import (
	"bufio"
	"encoding/json"
	"hippo-data-acquisition/commons/logger"
	"hippo-data-acquisition/commons/queue"
	"hippo-data-acquisition/config"
	"hippo-data-acquisition/outputs/output_collection"
	"os"
)

var (
	outputChan = make(chan queue.DataInfo, 5)
)

type OutputFile struct {
	filePath string
}

// InitPlugin 初始化参数
func (f *OutputFile) InitPlugin(config config.OutputConfig, pluginName string) {
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
func (f *OutputFile) ExeOutput(dataQueue queue.Queue) {
	go readDataInfo(dataQueue)
	go writeDataInfoToFile(f.filePath)
}

func readDataInfo(dataQueue queue.Queue) {
	for {
		if len(dataQueue.GetDataList()) > 0 {
			err, dataInfo := dataQueue.PopData()
			if err == nil {
				outputChan <- dataInfo
			}
		}
	}
}

func writeDataInfoToFile(filePath string) {
	outputFile, err := os.OpenFile(filePath, os.O_CREATE, 0666)
	if err != nil {
		logger.LogInfo("outputFile", "创建输入文件对象失败："+err.Error())
	}

	defer outputFile.Close()

	for {
		DataInfo, ok := <-outputChan
		if ok {
			strByte, err := json.Marshal(&DataInfo)
			if err != nil {
				logger.LogInfo("outputFile", "输出数据转换成json字符串失败！")
			}

			dataJson := string(strByte)

			logger.LogInfo("outputFile", "输出数据："+dataJson+"到"+outputFile.Name())
			writer := bufio.NewWriter(outputFile)
			writer.WriteString(dataJson + "\n")
			writer.Flush()

		}
		
	}

}

func init() {
	output_collection.Add("outputFile", &OutputFile{})
}
