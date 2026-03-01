package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) handleTraces(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(s.store.ListTraces())
}
