

## Operations and resilience endpoints

Each service exposes `GET /health/live`, `GET /health/ready`, and `GET /metrics`. Request logs include request ID, method, path, status, duration, and error. See `docs/resilience-and-operations.md` for Consul, round-robin load balancing, Product Service retries/circuit breaker, stock reservation/cancellation, and the 100-request concurrent load test.
