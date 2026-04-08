package monitor

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

type CpuMonitor struct {
}

func (m *CpuMonitor) Check(ctx context.Context) string {

	percent, err := cpu.PercentWithContext(ctx, 1*time.Second, false)
	if err != nil && len(percent) == 0 {
		return "N/A"
	}
	return fmt.Sprintf("Current CPU usage: %.2f%%", percent[0])
}
