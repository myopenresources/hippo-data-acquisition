package input_collection

import (
	"hippo-data-acquisition/commons/queue"
	"hippo-data-acquisition/config"
)

// InputPlugin 输入插件接口
type InputPlugin interface {
	// InitPlugin 初始化参数
	InitPlugin(config config.InputConfig)
	// PrepareCron 准备定时器
	PrepareCron()
	// BeforeExeDataAcquisition  执行数据采集前
	BeforeExeDataAcquisition()
	// ExeDataAcquisition  执行数据采集
	ExeDataAcquisition(dataQueue queue.Queue)
	// AfterExeDataAcquisition  执行数据采集后
	AfterExeDataAcquisition()
}

var (
	inputCollection = make(map[string]InputPlugin)
)

func Add(name string, input InputPlugin) {
	inputCollection[name] = input
}

func GetInputs() map[string]InputPlugin {
	return inputCollection
}
