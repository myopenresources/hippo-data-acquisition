package main

import (
	"hippo-data-acquisition/agent"
	"hippo-data-acquisition/config"
	_ "hippo-data-acquisition/inputs/register"
	_ "hippo-data-acquisition/outputs/register"
)

// main 主函数入口
func main() {
	// 初始化配置文件
	loadConfigRes := config.LoadConfig()

	if loadConfigRes {
		//初始化agent
		agent.InitAgent()
	}

	select {}
}
