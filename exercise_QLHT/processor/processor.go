package processor

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"

	"dwchwang.com/exercise_qlht/models"

	"github.com/shirou/gopsutil/v4/mem"
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
	var output string

	vmStat, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return fmt.Sprintf("[Get Top Processes] Could not retrieve process main info: %v \n ", err)
	}

	totalMemory := vmStat.Total

	processes, err := process.ProcessesWithContext(ctx)
	if err != nil {
		return fmt.Sprintf("[Get Top Processes] Could not retrieve process list: %v \n ", err)
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	var cpuList, memList []models.ProcStat

	procChan := make(chan models.ProcStat, len(processes))

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

				memInfo, err := p.MemoryInfoWithContext(ctx)
				if err != nil {
					return
				}

				ranPercent := (float64(memInfo.RSS) / float64(totalMemory)) * 100
				createTime, err := p.CreateTimeWithContext(ctx)
				if err != nil {
					return
				}

				runningTime := time.Since(time.Unix(createTime/1000, 0))

				if cpuPercent > 1 || ranPercent > 1 {
					mu.Lock()
					procStat := models.ProcStat{
						PID:           p.Pid,
						Name:          name,
						CPUPercent:    cpuPercent,
						MemoryUsed:    memInfo.RSS,
						MemoryPercent: ranPercent,
						RunningTime:   runningTime,
					}
					procChan <- procStat
					mu.Unlock()
				}
			}
		}(p)
	}

	go func() {
		wg.Wait()
		close(procChan)
	}()

	for stat := range procChan {
		if stat.CPUPercent > 5 {
			cpuList = append(cpuList, stat)
		}

		if stat.MemoryPercent > 5 {
			memList = append(memList, stat)
		}
	}

	sort.Slice(cpuList, func(i int, j int) bool {
		return cpuList[i].CPUPercent > cpuList[j].CPUPercent
	})

	sort.Slice(memList, func(i int, j int) bool {
		return memList[i].CPUPercent > memList[j].CPUPercent
	})

	output += "==== Top 5 CPU consuming processes ==== \n"
	for i := 0; i < len(cpuList) && i < 5; i++ {
		output += fmt.Sprintf("[%d] %s - CPU: %.2f%% - RAM: %.2f MB (%.2f%%) - Running Time: %s \n",
			cpuList[i].PID,
			cpuList[i].Name,
			cpuList[i].CPUPercent,
			float64(cpuList[i].MemoryUsed)/1024.0/1024.0,
			cpuList[i].MemoryPercent,
			cpuList[i].RunningTime,
		)
	}

	output += "==== Top 5 Ram consuming processes ==== \n"
	for i := 0; i < len(memList) && i < 5; i++ {
		output += fmt.Sprintf("[%d] %s - Mem: %.2f%% - RAM: %.2f MB (%.2f%%) - Running Time: %s \n",
			memList[i].PID,
			memList[i].Name,
			memList[i].CPUPercent,
			float64(memList[i].MemoryUsed)/1024.0/1024.0,
			memList[i].MemoryPercent,
			memList[i].RunningTime,
		)
	}

	return output
}
