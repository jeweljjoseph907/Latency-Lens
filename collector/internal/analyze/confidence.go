package analyze

import "latency-prism/collector/internal/model"

func ConfidenceScore(trace model.Trace, events []model.RuntimeEvent) float64 {
	if len(trace.Spans) == 0 {
		return 0
	}

	spanSignal := minFloat(float64(len(trace.Spans))/20.0, 1.0)
	eventSignal := minFloat(float64(len(events))/10.0, 1.0)

	return (spanSignal * 0.7) + (eventSignal * 0.3)
}

func minFloat(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
