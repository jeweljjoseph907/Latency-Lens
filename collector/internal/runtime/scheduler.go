package runtime

import (
	goruntime "runtime"
	"time"

	"latency-prism/collector/internal/model"
)

func CollectSchedulerEvent() model.RuntimeEvent {
	return model.RuntimeEvent{
		Timestamp: time.Now().UTC(),
		Type:      "scheduler",
		Metadata: map[string]any{
			"gomaxprocs": goruntime.GOMAXPROCS(0),
			"goroutines": goruntime.NumGoroutine(),
		},
	}
}
