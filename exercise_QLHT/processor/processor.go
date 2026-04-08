package processor

import (
	"context"
	"fmt"
	"sync"
	"time"

	"dwchwang.com/exercise_qlht/monitor"
)

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
			monitor := monitor.CpuMonitor{}
			result := monitor.Check(ctx)
			fmt.Println(result)
		}
	}
}