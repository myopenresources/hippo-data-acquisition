package config

import (
	"encoding/json"
	"fmt"
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
		fmt.Println("配置文件config读取失败！")
		return false
	}
	err = json.Unmarshal(Data, &DaqConfig)
	if err != nil {
		fmt.Println("配置文件config转换成对象失败:" + err.Error())
		return false
	}
	return true
}
