# Latency Prism Architecture

## Overview

Latency Prism is split into two Go modules:

- `app`: sample instrumented service generating latency traces.
- `collector`: ingestion, storage, analysis, and API serving of trace insights.

## Collector Flow

1. Ingestion adapters normalize incoming span payloads.
2. In-memory store groups spans by trace ID.
3. Analysis engine computes bottlenecks and component breakdowns.
4. HTTP API exposes `/traces` and `/breakdown` for consumers.
