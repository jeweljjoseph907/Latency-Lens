package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	APIAddr        string
	TraceRetention time.Duration
}

func LoadFromEnv() Config {
	addr := os.Getenv("COLLECTOR_ADDR")
	if addr == "" {
		addr = ":9090"
	}

	retentionSeconds := 600
	if raw := os.Getenv("TRACE_RETENTION_SECONDS"); raw != "" {
		if parsed, err := strconv.Atoi(raw); err == nil && parsed > 0 {
			retentionSeconds = parsed
		}
	}

	return Config{
		APIAddr:        addr,
		TraceRetention: time.Duration(retentionSeconds) * time.Second,
	}
}
