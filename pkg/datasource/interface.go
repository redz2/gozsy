package datasource

import (
	"context"
	"time"
)

type Querier interface {
	Query(ctx context.Context, QueryUrl string, QL string, ts time.Time) (result ResultType, err error)
	QueryRange(ctx context.Context, QueryUrl string, QL string, from time.Time, to time.Time, step time.Duration) (result ResultType, err error)
}

type Result interface {
	GetResult() interface{}
}

type ResultType int

const (
	ValNone ResultType = iota
	ValVictoriaMetrics
	ValLoki
)

type ResultValue []Sample

type Sample struct {
	Labels    []map[string]string
	TimeStamp time.Time
	Values    float64
}
