package main

import (
	"log"
	"net/http"
	"os"

	"latency-prism/app/internal/handlers"
	"latency-prism/app/internal/tracing"
)

func main() {
	shutdown, err := tracing.Setup("latency-prism-app")
	if err != nil {
		log.Fatalf("failed to initialize tracing: %v", err)
	}
	defer shutdown()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.Health)
	mux.HandleFunc("/work", handlers.Work)

	addr := os.Getenv("APP_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	log.Printf("sample app listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
