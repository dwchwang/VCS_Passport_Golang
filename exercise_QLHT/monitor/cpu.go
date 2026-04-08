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

func (m *CpuMonitor) Check(ctx context.Context) string {

	percent, err := cpu.PercentWithContext(ctx, 1*time.Second, false)
	if err != nil && len(percent) == 0 {
		return "N/A"
	}
	return fmt.Sprintf("%.2f%%", percent[0])
}
