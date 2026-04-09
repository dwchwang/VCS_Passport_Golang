package processor

import (
	"context"
	"fmt"
	"sync"
	"time"

	"dwchwang.com/exercise_qlht/models"

	"github.com/shirou/gopsutil/v4/process"
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

func GetTopProcessor(ctx context.Context) string {
	processes, err := process.ProcessesWithContext(ctx)
	if err != nil {
		return fmt.Sprintf("[Get Top Processes] Could not retrieve process list: %v \n ", err)
	}

	var wg sync.WaitGroup

	for _, p := range processes {
		wg.Add(1)
		go func(*process.Process) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				name, err := p.NameWithContext(ctx)
				if err != nil {
					return
				}
				
				cpuPercent, err := p.CPUPercentWithContext(ctx)
				if err != nil {
					return
				}

				memInfo, err := p.MemoryPercentWithContext(ctx)
				if err != nil {
					return 
				}

				fmt.Printf("%+v \n", name)
				fmt.Printf("%+v \n", cpuPercent)
				fmt.Printf("%+v \n", memInfo)
				fmt.Printf("=========================")
			}
		}(p)	
	}

	wg.Wait()
	return ""
}
