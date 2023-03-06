package logger

import (
	"fmt"
	"hippo-data-acquisition/commons/utils"
	"hippo-data-acquisition/config"
)

var (
	logChan     = make(chan string, 30)
	logFilePath = ""
)

func LogInfo(module string, msg string) {
	logJson := createLogJson(module, msg, "info")
	logChan <- logJson
	fmt.Println(logJson)
}

func LogError(module string, msg string) {
	logJson := createLogJson(module, msg, "error")
	logChan <- logJson
	fmt.Println(logJson)
}

func LogWarning(module string, msg string) {
	logJson := createLogJson(module, msg, "warn")
	logChan <- logJson
	fmt.Println(logJson)
}

func createLogJson(module string, msg string, logType string) string {
	return "{\"module\":" + module + ",\"msg\":\"" + msg + "\",\"logType\":\"" + logType + "\",\"createTime\":" + utils.GetNowTime("") + "\"}"
}

func WriteLogToFile() {
	for {
		log, ok := <-logChan
		if ok {
			// 检查是否已经超过一天了，超过重新创建名字
			checkLogFilePath := utils.NewDateFilePath(config.DaqConfig.LogPath, "")
			if checkLogFilePath != logFilePath {
				logFilePath = checkLogFilePath
			}

			utils.WriteStrToFile(logFilePath, log, "logger", func(log string) {
				fmt.Println(log)
			})
		}
	}
}

func InitLogger() {
	logPath := config.DaqConfig.LogPath
	if logPath != "" {
		go WriteLogToFile()
	}

}
