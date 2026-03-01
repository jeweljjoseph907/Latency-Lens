# Latency Prism Design Notes

## Principles

- Keep ingestion and analysis loosely coupled via domain models.
- Prefer simple deterministic heuristics for first iteration.
- Make advanced features (eBPF, critical path) optional and additive.

## Near-term Extensions

- Add OTLP/gRPC endpoint in `internal/ingestion`.
- Persist traces in Redis or ClickHouse.
- Add confidence calibration using historical traces.
