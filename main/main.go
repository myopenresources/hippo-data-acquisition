package main

import (
	"hippo-data-acquisition/agent"
	"hippo-data-acquisition/config"
)

// main 主函数入口
func main() {
	// 初始化配置文件
	config.LoadConfig()
	//初始化agent
	agent.InitAgent()
}
