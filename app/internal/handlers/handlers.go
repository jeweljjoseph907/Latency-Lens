package handlers

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"
	"time"
)

func Health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func Work(w http.ResponseWriter, _ *http.Request) {
	delay := time.Duration(20+rand.IntN(180)) * time.Millisecond
	time.Sleep(delay)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"simulated_latency_ms": delay.Milliseconds(),
	})
}
