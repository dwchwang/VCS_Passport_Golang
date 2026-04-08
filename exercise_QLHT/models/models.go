package models

import "context"

type Monitor interface {
	Name() string
	Check(ctx context.Context) string
}

type SystemStat struct {
	Name  string
	Value string
}
