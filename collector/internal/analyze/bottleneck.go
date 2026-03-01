package analyze

import (
	"time"

	"latency-prism/collector/internal/model"
)

func DetectBottleneck(trace model.Trace) (string, time.Duration) {
	var longest model.Span
	for _, span := range trace.Spans {
		if span.Duration() > longest.Duration() {
			longest = span
		}
	}
	return longest.Name, longest.Duration()
}
