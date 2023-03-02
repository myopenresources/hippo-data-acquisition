package mem

import (
	"github.com/shirou/gopsutil/v3/mem"
)

// GetMemPercent 获取内存使用率
func GetMemPercent() *mem.VirtualMemoryStat {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return nil
	}
	return memInfo
}
