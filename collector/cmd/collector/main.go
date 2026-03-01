package main

import (
	"log"

	"latency-prism/collector/internal/analyze"
	"latency-prism/collector/internal/api"
	"latency-prism/collector/internal/config"
	"latency-prism/collector/internal/store"
)

func main() {
	cfg := config.LoadFromEnv()

	spanStore := store.NewMemoryStore(cfg.TraceRetention)
	engine := analyze.NewEngine(spanStore)
	server := api.NewServer(spanStore, engine)

	log.Printf("collector API listening on %s", cfg.APIAddr)
	if err := server.Start(cfg.APIAddr); err != nil {
		log.Fatalf("collector server stopped: %v", err)
	}
}
