package agent

import (
	"fmt"
	"hippo-data-acquisition/commons/logger"
	"hippo-data-acquisition/commons/queue"
	"hippo-data-acquisition/config"
	"hippo-data-acquisition/plugins/inputs/fs_notify"
)

var (
	InputCronList []Cron = make([]Cron, 0)
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

		cron := Cron{
			spec: specVal.(string),
		}

		dataQueue := queue.NewDataQueue()
		cronOk := cron.Start(func() {
			// todo 反射获取input对象并且调用execute函数，将获取到的数据推送到处理器(私有处理，公共处理器)
			fs := fs_notify.FsNotify{}
			fs.ExeDataAcquisition(&dataQueue)

			for i2 := range dataQueue.GetDataList() {
				fmt.Println(dataQueue.GetDataList()[i2].Fields)
			}
			fmt.Println("================")
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
