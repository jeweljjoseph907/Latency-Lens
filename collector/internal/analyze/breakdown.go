package analyze

import (
	"strings"
	"time"

	"latency-prism/collector/internal/model"
)

func BuildBreakdown(trace model.Trace) []model.ComponentBreakdown {
	if len(trace.Spans) == 0 {
		return nil
	}

	total, _ := AggregateTrace(trace)
	if total <= 0 {
		return nil
	}

	byComponent := map[string]int64{}
	for _, span := range trace.Spans {
		component := componentFromSpanName(span.Name)
		byComponent[component] += span.Duration().Nanoseconds()
	}

	result := make([]model.ComponentBreakdown, 0, len(byComponent))
	for component, ns := range byComponent {
		duration := model.ComponentBreakdown{
			Component: component,
			Duration:  time.Duration(ns),
			Percent:   float64(ns) / float64(total.Nanoseconds()) * 100,
		}
		result = append(result, duration)
	}

	return result
}

func componentFromSpanName(name string) string {
	parts := strings.Split(name, ".")
	if len(parts) == 0 || parts[0] == "" {
		return "unknown"
	}
	return parts[0]
}
