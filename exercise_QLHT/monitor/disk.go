package monitor

import (
	"context"
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/v4/disk"
)

type DiskMonitor struct {
}

func(m *DiskMonitor) Name () string {
	return "Disk Monitor"
}

func (m *DiskMonitor) Check(ctx context.Context) (string, bool) {
	path := "/"
	if runtime.GOOS == "windows" {
		path = "C:\\"
	}
	diskStat, err := disk.UsageWithContext(ctx, path)
	if err != nil {
		return fmt.Sprintf("[Disk Monitor] Could not retrieve process list: %v \n ", err), false
	}
	return fmt.Sprintf("%.2f%%", diskStat.UsedPercent), diskStat.UsedPercent > 60
}
