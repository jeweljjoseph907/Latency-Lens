package analyze

import "latency-prism/collector/internal/model"

func CorrelateRuntimeEvents(trace model.Trace, events []model.RuntimeEvent) float64 {
	if len(trace.Spans) == 0 || len(events) == 0 {
		return 0
	}

	traceStart := trace.Spans[0].StartTime
	traceEnd := trace.Spans[0].EndTime
	for _, span := range trace.Spans[1:] {
		if span.StartTime.Before(traceStart) {
			traceStart = span.StartTime
		}
		if span.EndTime.After(traceEnd) {
			traceEnd = span.EndTime
		}
	}

	hits := 0
	for _, event := range events {
		if !event.Timestamp.Before(traceStart) && !event.Timestamp.After(traceEnd) {
			hits++
		}
	}

	return float64(hits) / float64(len(events))
}
