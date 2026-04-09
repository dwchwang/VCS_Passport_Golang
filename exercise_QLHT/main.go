package main

import (
	"context"
	"fmt"
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
		&monitor.NetMonitor{},
		&monitor.DiskMonitor{},
	}

	var wg sync.WaitGroup
	statCh := make(chan models.SystemStat)

	for _, m := range monitorList {
		wg.Add(1)
		go processor.RunMonitor(ctx, &wg, statCh, m)
	}

	go func() {
		for stat := range statCh {
			models.StatMutex.Lock()
			models.Stats[stat.Name] = stat
			models.StatMutex.Unlock()
		}
	}()
	printTicker := time.NewTicker(4 * time.Second)
	go func() {
		fmt.Println("===== System Status =====")
		for range printTicker.C {
			for _, stat := range models.Stats {
				models.StatMutex.Lock()
				fmt.Printf("[%s] %s \n", stat.Name, stat.Value)
				models.StatMutex.Unlock()
			}

			processor.GetTopProcessor(ctx)
		}
	}()

	time.Sleep(10 * time.Second)
	cancel()
	wg.Wait()
	close(statCh)
	printTicker.Stop()
}
