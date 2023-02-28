package output_collection

import (
	"hippo-data-acquisition/commons/queue"
	"hippo-data-acquisition/config"
)

// OutputPlugin 输出插件接口
type OutputPlugin interface {
	// InitPlugin 初始化参数
	InitPlugin(config config.OutputConfig, pluginName string)
	// BeforeExeOutput  执行输出前
	BeforeExeOutput()
	// ExeOutput  执行输出
	ExeOutput(dataQueue queue.Queue)
}

var (
	outputCollection = make(map[string]OutputPlugin)
)

func Add(name string, input OutputPlugin) {
	outputCollection[name] = input
}

func GetOutputs() map[string]OutputPlugin {
	return outputCollection
}
