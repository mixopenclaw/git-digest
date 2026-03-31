# Backend

This directory contains backend server code for the scanner API and related services.

Included in this patch series:

- POST /api/scan (queue a scan job)
- GET /api/scan/{id}/results (retrieve scan results)
- /metrics (Prometheus metrics)
- /health (readiness check)

Notes:
- Implementations are currently placeholders and intended as scaffolding for the next development pass.
- See backend/api/openapi.yaml for the documented endpoints.
