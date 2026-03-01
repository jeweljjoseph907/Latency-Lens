package store

import (
	"sync"
	"time"

	"latency-prism/collector/internal/model"
)

type MemoryStore struct {
	mu        sync.RWMutex
	traces    map[string]*model.Trace
	retention time.Duration
}

func NewMemoryStore(retention time.Duration) *MemoryStore {
	if retention <= 0 {
		retention = 10 * time.Minute
	}
	return &MemoryStore{
		traces:    make(map[string]*model.Trace),
		retention: retention,
	}
}

func (m *MemoryStore) AddSpan(span model.Span) {
	m.mu.Lock()
	defer m.mu.Unlock()

	trace, ok := m.traces[span.TraceID]
	if !ok {
		trace = model.NewTrace(span.TraceID)
		m.traces[span.TraceID] = trace
	}

	trace.AddSpan(span)
}

func (m *MemoryStore) GetTrace(traceID string) (*model.Trace, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	trace, ok := m.traces[traceID]
	if !ok {
		return nil, false
	}

	cp := *trace
	cp.Spans = append([]model.Span(nil), trace.Spans...)
	return &cp, true
}

func (m *MemoryStore) ListTraces() []model.Trace {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]model.Trace, 0, len(m.traces))
	for _, trace := range m.traces {
		cp := *trace
		cp.Spans = append([]model.Span(nil), trace.Spans...)
		result = append(result, cp)
	}

	return result
}
