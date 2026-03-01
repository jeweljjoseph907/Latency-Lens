package runtime

import (
	goruntime "runtime"
	"time"

	"latency-prism/collector/internal/model"
)

func CollectGCEvent() model.RuntimeEvent {
	var memStats goruntime.MemStats
	goruntime.ReadMemStats(&memStats)

	return model.RuntimeEvent{
		Timestamp: time.Now().UTC(),
		Type:      "gc",
		Metadata: map[string]any{
			"pause_total_ns": memStats.PauseTotalNs,
			"num_gc":         memStats.NumGC,
		},
	}
}
