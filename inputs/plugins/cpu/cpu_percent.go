package cpu

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"time"
)

// GetCPUPercent 获取CPU使用率
func GetCPUPercent() ([]cpu.InfoStat, []float64) {
	percents, _ := cpu.Percent(time.Second, false)
	cpuInfos, _ := cpu.Info()
	return cpuInfos, percents
}
