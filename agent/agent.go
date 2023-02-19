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
	runInputs()

}

func runInputs() {
	for i := range config.DaqConfig.Inputs {
		logger.LogInfo("agent", "正在启动input "+config.DaqConfig.Inputs[i].InputName+"！")
		params := config.DaqConfig.Inputs[i].Params

		specVal, ok := params["spec"]
		if !ok {
			specVal = config.DaqConfig.Agent.Spec
		}

		inputs := input_collection.GetInputs()
		plugin := inputs[config.DaqConfig.Inputs[i].InputName]

		// input插件不存在时
		if plugin == nil {
			logger.LogInfo("agent", "找不到input "+config.DaqConfig.Inputs[i].InputName+"！")
			continue
		}

		cron := Cron{
			spec: specVal.(string),
		}

		dataQueue := queue.NewDataQueue()
		cronOk := cron.Start(func() {
			plugin.ExeDataAcquisition(&dataQueue)

			//有消息时进行下一步
			if len(dataQueue.GetDataList()) > 0 {
				for i2 := range dataQueue.GetDataList() {
					fmt.Println(dataQueue.GetDataList()[i2].DataBody)
				}
				fmt.Println("================")
			} else {

				//如果没有消息释放当前定时器
				cron.Stop()

			}
		})

		if cronOk == nil {
			logger.LogInfo("agent", "启动input "+config.DaqConfig.Inputs[i].InputName+"成功！")
		}

		InputCronList = append(InputCronList, cron)
	}

}

func GetInputCronList() []Cron {
	return InputCronList
}
