package model

import "time"

type Span struct {
	TraceID      string
	SpanID       string
	ParentSpanID string
	Name         string
	StartTime    time.Time
	EndTime      time.Time
	Attributes   map[string]string
}

func (s Span) Duration() time.Duration {
	if s.EndTime.Before(s.StartTime) {
		return 0
	}
	return s.EndTime.Sub(s.StartTime)
}
