package ingestion

import (
	"fmt"
	"time"

	"latency-prism/collector/internal/model"
)

type SpanAdapter interface {
	Convert(raw any) (model.Span, error)
}

type MapAdapter struct{}

func (MapAdapter) Convert(raw any) (model.Span, error) {
	m, ok := raw.(map[string]any)
	if !ok {
		return model.Span{}, fmt.Errorf("unsupported payload type: %T", raw)
	}

	traceID, _ := m["trace_id"].(string)
	spanID, _ := m["span_id"].(string)
	name, _ := m["name"].(string)
	if traceID == "" || spanID == "" || name == "" {
		return model.Span{}, fmt.Errorf("trace_id, span_id and name are required")
	}

	now := time.Now().UTC()
	return model.Span{
		TraceID:    traceID,
		SpanID:     spanID,
		Name:       name,
		StartTime:  now,
		EndTime:    now.Add(1 * time.Millisecond),
		Attributes: map[string]string{},
	}, nil
}
