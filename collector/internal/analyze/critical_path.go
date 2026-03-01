package analyze

import (
	"sort"

	"latency-prism/collector/internal/model"
)

func CriticalPath(trace model.Trace) []model.Span {
	if len(trace.Spans) == 0 {
		return nil
	}

	spans := append([]model.Span(nil), trace.Spans...)
	sort.Slice(spans, func(i, j int) bool {
		return spans[i].Duration() > spans[j].Duration()
	})

	if len(spans) > 3 {
		return spans[:3]
	}
	return spans
}
