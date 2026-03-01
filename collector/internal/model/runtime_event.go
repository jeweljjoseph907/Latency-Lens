package model

import "time"

type RuntimeEvent struct {
	Timestamp time.Time
	Type      string
	Duration  time.Duration
	Metadata  map[string]any
}
