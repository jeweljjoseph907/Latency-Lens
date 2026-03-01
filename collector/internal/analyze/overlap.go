package analyze

import (
	"time"

	"latency-prism/collector/internal/model"
)

func ComputeOverlap(trace model.Trace) time.Duration {
	var overlap time.Duration
	for i := 0; i < len(trace.Spans); i++ {
		for j := i + 1; j < len(trace.Spans); j++ {
			start := maxTime(trace.Spans[i].StartTime, trace.Spans[j].StartTime)
			end := minTime(trace.Spans[i].EndTime, trace.Spans[j].EndTime)
			if end.After(start) {
				overlap += end.Sub(start)
			}
		}
	}
	return overlap
}

func maxTime(a, b time.Time) time.Time {
	if a.After(b) {
		return a
	}
	return b
}

func minTime(a, b time.Time) time.Time {
	if a.Before(b) {
		return a
	}
	return b
}
