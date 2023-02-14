package agent

import (
	"fmt"
	"hippo-data-acquisition/config"
	"time"
)

var (
	InputCronList []Cron = make([]Cron, 0)
)

func InitAgent() {
	runInputs()

}

func runInputs() {
	for i := range config.DaqConfig.Inputs {
		fmt.Println("正在启动input " + config.DaqConfig.Inputs[i].InputName + "！")
		params := config.DaqConfig.Inputs[i].Params

		specVal, ok := params["spec"]
		if !ok {
			specVal = config.DaqConfig.Agent.Spec
		}

		cron := Cron{
			spec: specVal.(string),
		}
		inputCron := cron.Start(func() {
			fmt.Println(time.Now())
			// todo 反射获取input对象并且调用execute函数，将获取到的数据推送到处理器(私有处理，公共处理器)
		})

		if inputCron == nil {
			fmt.Println("启动input " + config.DaqConfig.Inputs[i].InputName + "成功！")
		}

		InputCronList = append(InputCronList, cron)
	}

}

func GetInputCronList() []Cron {
	return InputCronList
}
