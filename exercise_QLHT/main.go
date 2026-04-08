package main

import (
	"context"
	"sync"
	"time"

	"dwchwang.com/exercise_qlht/processor"
)





func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go processor.RunMonitor(ctx, &wg)

	time.Sleep(10 * time.Second)
	cancel()
	wg.Wait()
}
