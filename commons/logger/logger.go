package logger

import (
	"fmt"
	"hippo-data-acquisition/commons/utils"
	"hippo-data-acquisition/config"
)

var (
	logChan     = make(chan string, 50)
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
			utils.WriteStrToFile(logFilePath, log, "logger", func(log string) {
				fmt.Println(log)
			})
		}
	}
}

func InitLogger() {
	logPath := config.DaqConfig.LogPath
	if logPath != "" {
		logFilePath = utils.NewDateFilePath(logPath, "")
		go WriteLogToFile()
	}

}
