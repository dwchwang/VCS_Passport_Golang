package monitor

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

type CpuMonitor struct {
}

func(m *CpuMonitor) Name () string {
	return "CPU Monitor"
}

func (m *CpuMonitor) Check(ctx context.Context) (string, bool) {

	cpuStat, err := cpu.PercentWithContext(ctx, 1*time.Second, false)
	if err != nil && len(cpuStat) == 0 {
		return fmt.Sprintf("[CPU Monitor] Could not retrieve process list: %v \n ", err), false
	}
	return fmt.Sprintf("%.2f%%", cpuStat[0]), cpuStat[0] > 60
}
