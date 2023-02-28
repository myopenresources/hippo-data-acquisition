package agent

import (
	"golang.org/x/exp/maps"
	"hippo-data-acquisition/commons/logger"
	"hippo-data-acquisition/commons/queue"
	"hippo-data-acquisition/config"
	"hippo-data-acquisition/inputs/input_collection"
	"hippo-data-acquisition/outputs/output_collection"
	"hippo-data-acquisition/processors/processors_collection"
)

var (
	InputCronList = make([]Cron, 0)
)

func InitAgent() {
	globalTag := config.DaqConfig.Tag

	if globalTag == nil {
		globalTag = make(map[string]string)
	}
	inputs := input_collection.GetInputs()

	if len(inputs) <= 0 {
		logger.LogInfo("agent", "未配置输入插件！")
		return
	}

	for i := range config.DaqConfig.Inputs {
		inputConfig := config.DaqConfig.Inputs[i]
		logger.LogInfo("agent", "正在启动input "+inputConfig.InputName+"！")
		params := config.DaqConfig.Inputs[i].Params

		//插件是否配置定时器表达式，如配置则使用插件配置，否则使用全局配置
		specVal, ok := params["spec"]
		if !ok {
			specVal = config.DaqConfig.Agent.Spec
		}

		//根据配置的插件名获取插件对象
		plugin := inputs[inputConfig.InputName]
		inputTag := inputConfig.Tag

		//将全局tag与插件的tag合并
		if inputTag != nil {
			maps.Copy(globalTag, inputTag)
		}

		// input插件不存在时
		if plugin == nil {
			logger.LogInfo("agent", "找不到input "+inputConfig.InputName+"！")
			continue
		}

		//初始化参数
		plugin.InitPlugin(inputConfig)

		//初始化定时器
		cron := Cron{
			spec: specVal.(string),
		}

		//准备定时器
		plugin.PrepareCron()

		dataQueue := queue.NewDataQueue(globalTag)
		cronOk := cron.Start(func() {
			//运行插件
			inputData, dataQueue := runInput(plugin, dataQueue)
			if inputData {
				inputProcessorData, dataQueue := runInputProcessors(inputConfig, inputConfig.InputName, dataQueue)
				if inputProcessorData {
					processorData, dataQueue := runCommonProcessors(config.DaqConfig.Processors, inputConfig.InputName, dataQueue)
					if processorData {
						runOutPuts(config.DaqConfig.Outputs, inputConfig.InputName, dataQueue)
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

func runInput(plugin input_collection.InputPlugin, dataQueue queue.DataQueue) (bool, queue.DataQueue) {
	plugin.BeforeExeDataAcquisition()
	plugin.ExeDataAcquisition(&dataQueue)
	plugin.AfterExeDataAcquisition()

	//有消息时进行下一步
	if len(dataQueue.GetDataList()) > 0 {
		return true, dataQueue
	}

	return false, dataQueue
}

// runInputProcessors 运行input插件的处理插件
func runInputProcessors(inputConfig config.InputConfig, pluginName string, dataQueue queue.DataQueue) (bool, queue.DataQueue) {
	return runProcessors(inputConfig.Processors, pluginName, dataQueue, false)
}

// runProcessors 运行公共处理插件
func runCommonProcessors(processorsConfig []config.ProcessorConfig, pluginName string, dataQueue queue.DataQueue) (bool, queue.DataQueue) {
	return runProcessors(processorsConfig, pluginName, dataQueue, true)
}

// runProcessor 运行处理器插件
func runProcessors(processorsConfig []config.ProcessorConfig, pluginName string, dataQueue queue.DataQueue, isCommonProcessor bool) (bool, queue.DataQueue) {
	processors := processors_collection.GetProcessors()
	if len(processors) <= 0 {
		logger.LogInfo("agent", "未配置处理插件！")
		return true, dataQueue
	}

	for i := range processorsConfig {
		processorConfig := processorsConfig[i]

		processor := processors[processorConfig.ProcessorsName]

		if processor == nil {
			logger.LogInfo("agent", "找不到处理插件 "+processorConfig.ProcessorsName+"！")
			continue
		}

		processor.InitPlugin(processorConfig, pluginName)

		processor.BeforeExeProcess()
		processor.ExeProcess(&dataQueue)
		processor.AfterExeProcess()

		if len(dataQueue.GetDataList()) <= 0 {
			if isCommonProcessor {
				logger.LogInfo("agent", pluginName+"对应的的处理器"+processorConfig.ProcessorsName+"返回的数据为空！")
			} else {
				logger.LogInfo("agent", "公共处理器"+processorConfig.ProcessorsName+"返回的数据为空！")
			}

			return false, dataQueue
		}

	}
	return true, dataQueue
}

func runOutPuts(outputsConfig []config.OutputConfig, pluginName string, dataQueue queue.DataQueue) {
	outputs := output_collection.GetOutputs()
	if len(outputs) <= 0 {
		logger.LogInfo("agent", "未配置输出插件！")
		return
	}
	for i := range outputsConfig {
		outputConfig := outputsConfig[i]

		output := outputs[outputConfig.OutputName]
		if output == nil {
			logger.LogInfo("agent", "找不到输出插件 "+outputConfig.OutputName+"！")
			continue
		}

		output.InitPlugin(outputConfig, pluginName)
		output.BeforeExeOutput()
		output.ExeOutput(&dataQueue)

	}

}

func GetInputCronList() []Cron {
	return InputCronList
}
