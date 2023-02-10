package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	// DataAcquisitionConfig 采集器配置
	DataAcquisitionConfig Config
)

// LoadConfig 加载配置
func LoadConfig() {
	wd, _ := os.Getwd()
	Data, err := os.ReadFile(wd + "/config/config.json")
	if err != nil {
		fmt.Println("配置文件config读取失败！")
	}
	err = json.Unmarshal(Data, &DataAcquisitionConfig)
	if err != nil {
		fmt.Println("配置文件config转换成对象失败:" + err.Error())
	}
}
