package api

import (
	"log"
	"net/http"

	"latency-prism/collector/internal/analyze"
	"latency-prism/collector/internal/store"
)

type Server struct {
	store  store.Store
	engine *analyze.Engine
	mux    *http.ServeMux
}

func NewServer(store store.Store, engine *analyze.Engine) *Server {
	s := &Server{
		store:  store,
		engine: engine,
		mux:    http.NewServeMux(),
	}

	s.routes()
	return s
}

func (s *Server) routes() {
	s.mux.HandleFunc("/health", s.handleHealth)
	s.mux.HandleFunc("/traces", s.handleTraces)
	s.mux.HandleFunc("/breakdown", s.handleBreakdown)
}

func (s *Server) Start(addr string) error {
	log.Printf("http API routes: /health, /traces, /breakdown")
	return http.ListenAndServe(addr, s.mux)
}

func (s *Server) handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}
