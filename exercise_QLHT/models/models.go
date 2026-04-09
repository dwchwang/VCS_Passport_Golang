package models

import (
	"context"
	"sync"
	"time"
)

type Monitor interface {
	Name() string
	Check(ctx context.Context) string
}

type SystemStat struct {
	Name  string
	Value string
}

type ProcStat struct {
	PID           int32
	Name          string
	CPUPercent    float64
	MemoryUsed    uint64
	MemoryPercent float64
	RunningTime   time.Duration
}

var (
	StatMutex sync.Mutex
	Stats     = map[string]SystemStat{}
)
