package monitor

import (
	"context"
	"fmt"

	"github.com/shirou/gopsutil/v4/mem"
)

type MemMonitor struct {
}

func(m *MemMonitor) Name () string {
	return "Memory Monitor"
}

func (m *MemMonitor) Check(ctx context.Context) (string, bool) {

	vmStat, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return fmt.Sprintf("[Mem Monitor] Could not retrieve process list: %v \n ", err), false
	}
	return fmt.Sprintf("%.2f%%", vmStat.UsedPercent), vmStat.UsedPercent > 60
}
