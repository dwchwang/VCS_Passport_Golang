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

func (m *MemMonitor) Check(ctx context.Context) string {

	vmStat, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return "N/A"
	}
	return fmt.Sprintf("%.2f%%", vmStat.UsedPercent)
}
