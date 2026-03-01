package analyze

import (
	"fmt"
	"time"

	"latency-prism/collector/internal/model"
	"latency-prism/collector/internal/store"
)

type Engine struct {
	store store.Store
}

func NewEngine(store store.Store) *Engine {
	return &Engine{store: store}
}

func (e *Engine) TraceBreakdown(traceID string) ([]model.ComponentBreakdown, error) {
	trace, ok := e.store.GetTrace(traceID)
	if !ok {
		return nil, fmt.Errorf("trace not found: %s", traceID)
	}

	return BuildBreakdown(*trace), nil
}

func (e *Engine) Bottleneck(traceID string) (string, time.Duration, error) {
	trace, ok := e.store.GetTrace(traceID)
	if !ok {
		return "", 0, fmt.Errorf("trace not found: %s", traceID)
	}

	name, duration := DetectBottleneck(*trace)
	return name, duration, nil
}
