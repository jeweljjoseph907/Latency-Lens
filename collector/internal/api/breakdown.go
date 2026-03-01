package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) handleBreakdown(w http.ResponseWriter, r *http.Request) {
	traceID := r.URL.Query().Get("trace_id")
	if traceID == "" {
		http.Error(w, "missing trace_id", http.StatusBadRequest)
		return
	}

	breakdown, err := s.engine.TraceBreakdown(traceID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(breakdown)
}
