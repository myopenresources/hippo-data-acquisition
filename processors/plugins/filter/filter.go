package filter

import (
	"fmt"
	"hippo-data-acquisition/commons/queue"
	"hippo-data-acquisition/config"
	"hippo-data-acquisition/processors/processors_collection"
)

type Filter struct {
}

// InitPlugin 初始化参数
func (f *Filter) InitPlugin(config config.ProcessorConfig, pluginName string) {
	fmt.Println(config)
	fmt.Println(pluginName)
}

// BeforeExeProcess  执行处理前
func (f *Filter) BeforeExeProcess() {
	fmt.Println("BeforeExeProcess")
}

// ExeProcess  执行处理
func (f *Filter) ExeProcess(dataQueue queue.Queue) {
	fmt.Println("ExeProcess")
	fmt.Println(dataQueue)
	//dataQueue.SetDataList(make([]queue.DataInfo, 0))
}

// AfterExeProcess  执行处理后
func (f *Filter) AfterExeProcess() {
	fmt.Println("AfterExeProcess")
}

func init() {
	processors_collection.Add("filter", &Filter{})
}
