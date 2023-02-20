package agent

import (
	"fmt"
	"hippo-data-acquisition/commons/logger"
	"hippo-data-acquisition/commons/queue"
	"hippo-data-acquisition/config"
	"hippo-data-acquisition/inputs/input_collection"
)

var (
	InputCronList = make([]Cron, 0)
)

func InitAgent() {
	for i := range config.DaqConfig.Inputs {
		inputConfig := config.DaqConfig.Inputs[i]
		logger.LogInfo("agent", "正在启动input "+inputConfig.InputName+"！")
		params := config.DaqConfig.Inputs[i].Params

		specVal, ok := params["spec"]
		if !ok {
			specVal = config.DaqConfig.Agent.Spec
		}

		inputs := input_collection.GetInputs()
		plugin := inputs[inputConfig.InputName]

		// input插件不存在时
		if plugin == nil {
			logger.LogInfo("agent", "找不到input "+inputConfig.InputName+"！")
			continue
		}

		//初始化参数
		plugin.InitPlugin(inputConfig)

		cron := Cron{
			spec: specVal.(string),
		}

		//准备定时器
		plugin.PrepareCron()

		dataQueue := queue.NewDataQueue()
		cronOk := cron.Start(func() {
			//运行插件
			inputData := runInput(plugin, dataQueue)
			if inputData {
				inputProcessorData := runInputProcessors(inputConfig, plugin, dataQueue)
				if inputProcessorData {
					processorData := runProcessors(inputConfig, plugin, dataQueue)
					if processorData {
						runOutPuts()
					}
				}
			}
		})

		if cronOk == nil {
			logger.LogInfo("agent", "启动input "+config.DaqConfig.Inputs[i].InputName+"成功！")
		}

		InputCronList = append(InputCronList, cron)
	}

}

func runInput(plugin input_collection.InputPlugin, dataQueue queue.DataQueue) bool {
	plugin.BeforeExeDataAcquisition()
	plugin.ExeDataAcquisition(&dataQueue)
	plugin.AfterExeDataAcquisition()

	//有消息时进行下一步
	if len(dataQueue.GetDataList()) > 0 {
		return true
	}

	return false
}

func runInputProcessors(config config.InputConfig, plugin input_collection.InputPlugin, dataQueue queue.DataQueue) bool {
	for i := range config.Processors {
		processor := config.Processors[i]
		fmt.Println(processor.ProcessorsName)
	}
	return true
}

func runProcessors(config config.InputConfig, plugin input_collection.InputPlugin, dataQueue queue.DataQueue) bool {
	return true
}

func runOutPuts() {

}

func GetInputCronList() []Cron {
	return InputCronList
}
