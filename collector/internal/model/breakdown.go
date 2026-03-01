package model

import "time"

type ComponentBreakdown struct {
	Component string        `json:"component"`
	Duration  time.Duration `json:"duration"`
	Percent   float64       `json:"percent"`
}
