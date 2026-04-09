package processor

import (
	"context"
	"fmt"
	"os"
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
			value, alert := m.Check(ctx)
			stat :=  models.SystemStat{
				Name:  m.Name(),
				Value: value,
			}
			statCh <- stat

			if alert {
				LogAlert(stat)
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
				if cpuPercent < 1 || ranPercent < 1 {
					return
				}

				createTime, err := p.CreateTimeWithContext(ctx)
				if err != nil {
					return
				}

				runningTime := time.Since(time.Unix(createTime/1000, 0))

				procStat := models.ProcStat{
					PID:           p.Pid,
					Name:          name,
					CPUPercent:    cpuPercent,
					MemoryUsed:    memInfo.RSS,
					MemoryPercent: ranPercent,
					RunningTime:   runningTime,
				}
				procChan <- procStat

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
		return memList[i].MemoryPercent > memList[j].MemoryPercent
	})

	output += "==== Top 5 CPU consuming processes ==== \n"
	topCPU := 5
	if len(cpuList) < topCPU {
		topCPU = len(cpuList)
	}
	for key, c := range cpuList[:topCPU] {
		output += fmt.Sprintf("%d. [%d] %s - CPU: %.2f%% - RAM: %.2f MB (%.2f%%) - Running Time: %s \n",
			key+1,
			c.PID,
			c.Name,
			c.CPUPercent,
			float64(c.MemoryUsed)/1024.0/1024.0,
			c.MemoryPercent,
			c.RunningTime,
		)
	}

	topRam := 5
	if len(memList) < topRam {
		topRam = len(memList)
	}
	output += "==== Top 5 Ram consuming processes ==== \n"
	for key, r := range memList[:topRam] {
		output += fmt.Sprintf("%d. [%d] %s - Mem: %.2f%% - RAM: %.2f MB (%.2f%%) - Running Time: %s \n",
			key+1,
			r.PID,
			r.Name,
			r.CPUPercent,
			float64(r.MemoryUsed)/1024.0/1024.0,
			r.MemoryPercent,
			r.RunningTime,
		)
	}

	ExportToCSV(cpuList, memList)
	return output
}

func ExportToCSV(cpuList, memList []models.ProcStat) {
	f, err := os.OpenFile("process_stats.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Sprintf("[Export CSV Error] Fail to export csv: %v \n ", err)
		return
	}

	defer f.Close()

	if stat, err := f.Stat(); err == nil && stat.Size() == 0 {
		f.WriteString("TimeStamp,PID,Name,CPU (%),Memory (MB),Ram (%),RunningTime \n")
	}

	topCPU := 5
	if len(cpuList) < topCPU {
		topCPU = len(cpuList)
	}
	for _, c := range cpuList[:topCPU] {
		line := fmt.Sprintf("%s,%d,%s,%.2f,%.2f,%.2f,%s \n",
			time.Now().Format(time.RFC3339),
			c.PID,
			c.Name,
			c.CPUPercent,
			float64(c.MemoryUsed)/1024.0/1024.0,
			c.MemoryPercent,
			c.RunningTime,
		)
		f.WriteString(line)
	}

	topRam := 5
	if len(memList) < topRam {
		topRam = len(memList)
	}
	for _, r := range memList[:topRam] {
		line := fmt.Sprintf("%s,%d,%s,%.2f,%.2f,%.2f,%s \n",
			time.Now().Format(time.RFC3339),
			r.PID,
			r.Name,
			r.CPUPercent,
			float64(r.MemoryUsed)/1024.0/1024.0,
			r.MemoryPercent,
			r.RunningTime,
		)
		f.WriteString(line)
	}
}


func LogAlert(stat models.SystemStat) {
	models.StatMutex.Lock()
	defer models.StatMutex.Unlock()
	f, err := os.OpenFile("alert.log.", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Sprintf("[Log Alert] Fail to alert log: %v \n ", err)
		return
	}

	defer f.Close()

	logLine := fmt.Sprintf("[%s] ALERT: %s = %s \n", 
		time.Now().Format(time.RFC3339),
		stat.Name,
		stat.Value,
	)

	f.WriteString(logLine)
}