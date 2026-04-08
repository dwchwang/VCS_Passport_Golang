package main

import (
	"context"
	"sync"
	"time"

	"dwchwang.com/exercise_qlht/models"
	"dwchwang.com/exercise_qlht/monitor"
	"dwchwang.com/exercise_qlht/processor"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	monitorList := []models.Monitor{
		&monitor.CpuMonitor{},
		&monitor.MemMonitor{},
	}

	var wg sync.WaitGroup
	statCh := make(chan models.SystemStat, 2)

	for _, m := range monitorList {
		wg.Add(1)
		go processor.RunMonitor(ctx, &wg, statCh, m)
	}

	go func() {
		for stat := range statCh {
			println(stat.Name + ": " + stat.Value)
		}
	}()

	time.Sleep(10 * time.Second)
	cancel()
	wg.Wait()
	close(statCh)
}
