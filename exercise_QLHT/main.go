package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

type CpuMonitor struct {
}

func (m *CpuMonitor) Check(ctx context.Context) string {

	percent, err := cpu.PercentWithContext(ctx, 1*time.Second, false)
	if err != nil {
		return "Error getting CPU percent: " + err.Error()
	}
	return fmt.Sprintf("Current CPU usage: %.2f%%", percent[0])
}

func RunMonitor(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("CPU monitor stopped.")
			return
		case <-ticker.C:
			monitor := &CpuMonitor{}
			result := monitor.Check(ctx)
			fmt.Println(result)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go RunMonitor(ctx, &wg)

	wg.Wait()
	cancel()
}
