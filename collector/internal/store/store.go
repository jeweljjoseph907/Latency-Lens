package store

import (
	"time"

	"latency-prism/collector/internal/model"
)

type Store interface {
	AddSpan(span model.Span)
	GetTrace(traceID string) (*model.Trace, bool)
	ListTraces() []model.Trace
	Cleanup(now time.Time)
}
