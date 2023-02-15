package config

import (
	"encoding/json"
	"hippo-data-acquisition/commons/logger"
	"os"
)

var (
	// DaqConfig 采集器配置
	DaqConfig DataAcquisitionConfig
)

// LoadConfig 加载配置
func LoadConfig() bool {
	wd, _ := os.Getwd()
	Data, err := os.ReadFile(wd + "/config/config.json")
	if err != nil {
		logger.LogError("config", "配置文件config读取失败！")
		return false
	}
	err = json.Unmarshal(Data, &DaqConfig)
	if err != nil {
		logger.LogError("config", "配置文件config转换成对象失败:"+err.Error())
		return false
	}
	return true
}
