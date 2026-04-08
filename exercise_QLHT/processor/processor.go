package processor

import (
	"context"
	"fmt"
	"sync"
	"time"

	"dwchwang.com/exercise_qlht/models"
)

func RunMonitor(ctx context.Context, wg *sync.WaitGroup, statCh chan<- models.SystemStat, m models.Monitor) {
	defer wg.Done()
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("CPU monitor stopped.")
			return
		case <-ticker.C:
			statCh <- models.SystemStat{
				Name:  m.Name(),
				Value: m.Check(ctx),
			}
		}
	}
}
