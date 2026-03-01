package store

import "time"

func (m *MemoryStore) Cleanup(now time.Time) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for id, trace := range m.traces {
		if now.Sub(trace.UpdatedAt) > m.retention {
			delete(m.traces, id)
		}
	}
}
