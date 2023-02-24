package processors_collection

import (
	"hippo-data-acquisition/commons/queue"
	"hippo-data-acquisition/config"
)

// ProcessorPlugin 输入插件接口
type ProcessorPlugin interface {
	// InitPlugin 初始化参数
	InitPlugin(config config.ProcessorConfig, pluginName string)
	// BeforeExeProcess  执行处理前
	BeforeExeProcess()
	// ExeProcess  执行处理
	ExeProcess(dataQueue queue.Queue)
	// AfterExeProcess  执行处理后
	AfterExeProcess()
}

var (
	processorCollection = make(map[string]ProcessorPlugin)
)

func Add(name string, processor ProcessorPlugin) {
	processorCollection[name] = processor
}

func GetProcessors() map[string]ProcessorPlugin {
	return processorCollection
}
