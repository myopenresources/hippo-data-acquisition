package filter

import (
	"hippo-data-acquisition/commons/queue"
	"hippo-data-acquisition/config"
	"hippo-data-acquisition/processors/processors_collection"
)

type Filter struct {
}

// InitPlugin 初始化参数
func (f *Filter) InitPlugin(config config.ProcessorConfig, pluginName string) {

}

// BeforeExeProcess  执行处理前
func (f *Filter) BeforeExeProcess() {

}

// ExeProcess  执行处理
func (f *Filter) ExeProcess(dataQueue queue.Queue) {

}

// AfterExeProcess  执行处理后
func (f *Filter) AfterExeProcess() {

}

func init() {
	processors_collection.Add("filter", &Filter{})
}
