package model

import "time"

type Trace struct {
	ID        string
	Spans     []Span
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTrace(id string) *Trace {
	now := time.Now().UTC()
	return &Trace{ID: id, CreatedAt: now, UpdatedAt: now}
}

func (t *Trace) AddSpan(span Span) {
	t.Spans = append(t.Spans, span)
	t.UpdatedAt = time.Now().UTC()
}
