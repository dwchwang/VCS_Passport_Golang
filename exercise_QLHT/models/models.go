package models

import (
	"context"
	"sync"
)

type Monitor interface {
	Name() string
	Check(ctx context.Context) string
}

type SystemStat struct {
	Name  string
	Value string
}

var (
	StatMutex sync.Mutex
	Stats = map[string]SystemStat{}
)
