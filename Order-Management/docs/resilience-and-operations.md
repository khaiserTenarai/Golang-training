# Resilience, observability, load balancing, and inventory

## New endpoints

Each Go service exposes:

- `GET /health/live` — process liveness.
- `GET /health/ready` — readiness; Product and Order services ping PostgreSQL.
- `GET /metrics` — Prometheus text metrics: HTTP request count, errors, latency histogram/sum/count, and database errors.
- Existing `GET /health` remains available for compatibility.

Request logs are structured `slog` records with `request_id`, `method`, `path`, `status`, `duration`, and `error`. Clients may pass `X-Request-ID`; otherwise the service generates one.

## Consul and load balancing

Services register themselves in Consul at startup with IDs such as `product-service-8081`. The Gateway resolves healthy instances from Consul and uses round-robin selection with a short cache. If Consul is unavailable, it uses the configured `*_SERVICE_URL` fallback.

Start Consul natively on Windows (no Docker) by running ` .\scripts\start-consul.ps1` from the project root. Keep that terminal open and visit `http://localhost:8500`. Set `SERVICE_ADVERTISE_ADDRESS=127.0.0.1` when Consul and the Go services run on the same machine.

## Product-call resilience

Order Service calls Product Service using a request context and a per-attempt timeout. It retries transient errors up to three attempts with 100ms and 200ms exponential backoff. A simple circuit breaker opens after five failed calls and stays open for ten seconds before allowing a trial call. Cancellation of the request context interrupts the call or retry delay.

## Stock consistency

Product Service owns the product stock. The internal stock-reservation API uses a database transaction, a row lock, and a stock-operation ledger with a unique `(order_reference, action)` index. Stock is deducted only when enough stock remains. Repeated reserve/release calls using the same reference are idempotent.

Order Service reserves stock before persisting the order. If order persistence fails, it attempts a compensating release. Cancellation changes the order state, releases stock, and records `stockReleased`; if the release call fails, repeat the cancellation request to retry it safely. This is a saga/compensation approach, not a distributed ACID transaction; persistent failures should be monitored and reconciled.

## 100 concurrent order requests

Make sure Product Service is running and create a product with enough stock for the test. Then run from the project root:

```powershell
.\scripts\load-test.ps1 -ProductId 1 -CustomerId 1 -Quantity 1 -Requests 100
```

Or run the Go load tester directly:

```powershell
go run ./cmd/load-test -url http://localhost:8082 -product-id 1 -customer-id 1 -quantity 1 -requests 100
```

If stock is less than the total requested quantity, successful orders plus stock-conflict responses are expected. A stock conflict is not a server failure; it demonstrates that concurrent requests did not oversell inventory.

## Manual verification

```powershell
Invoke-RestMethod http://localhost:8080/health/live
Invoke-RestMethod http://localhost:8080/health/ready
Invoke-WebRequest http://localhost:8080/metrics
Invoke-RestMethod http://localhost:8500/v1/health/service/product-service?passing=true
```
