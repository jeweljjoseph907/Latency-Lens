package ingestion

import (
	"fmt"

	"latency-prism/collector/internal/store"
)

type OTLPServer struct {
	store   store.Store
	adapter SpanAdapter
}

func NewOTLPServer(store store.Store, adapter SpanAdapter) *OTLPServer {
	if adapter == nil {
		adapter = MapAdapter{}
	}
	return &OTLPServer{store: store, adapter: adapter}
}

func (s *OTLPServer) Ingest(raw any) error {
	span, err := s.adapter.Convert(raw)
	if err != nil {
		return fmt.Errorf("convert incoming payload: %w", err)
	}

	s.store.AddSpan(span)
	return nil
}
