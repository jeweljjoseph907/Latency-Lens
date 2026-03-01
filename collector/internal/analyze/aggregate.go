package analyze

import (
	"time"

	"latency-prism/collector/internal/model"
)

func AggregateTrace(trace model.Trace) (total time.Duration, spanCount int) {
	for _, span := range trace.Spans {
		total += span.Duration()
	}
	return total, len(trace.Spans)
}
