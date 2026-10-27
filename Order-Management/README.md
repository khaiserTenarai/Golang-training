# Order Management System (Go) — Local Windows Setup (No Docker)

This project contains Product Service, Order Service, Auth Service, and API Gateway using Gin, GraphQL, GORM/PostgreSQL, Swagger, JWT, and Consul. Docker is not required. PostgreSQL and Consul must run locally as native applications/services.

## Prerequisites

- Go installed and available in PATH
- PostgreSQL installed and running; create the database configured by `DB_NAME`
- Consul Windows binary (`consul.exe`) installed and available in PATH

## Run locally without Docker

1. Copy `.env.example` to `.env` if needed. Set PostgreSQL credentials, a strong `JWT_SECRET`, and the same `INTERNAL_SERVICE_TOKEN` for services that use internal service authentication.
2. Start PostgreSQL and create the database named in `DB_NAME`.
3. Edit `start-all.ps1` and set `$ConsulExe` to the location of `consul.exe` if Consul is not in PATH. Run the script from this project folder. It starts Consul if needed, builds the Go packages, and launches each service in its own PowerShell window. Consul UI: `http://localhost:8500`. Development mode is for local testing only.
4. Run `setup-consul-config.ps1` to store and verify the sample service URLs and resilience settings in Consul KV.
5. The local `.env` uses `SERVICE_ADVERTISE_ADDRESS=127.0.0.1`, appropriate when Consul and the Go services run on the same Windows machine. The Go services register themselves with Consul.

## Health, metrics, and Swagger

- Gateway: `http://localhost:8080/health/live`, `/health/ready`, `/metrics`
- Product Service: `http://localhost:8081/health/live`, `/health/ready`, `/metrics`
- Order Service: `http://localhost:8082/health/live`, `/health/ready`, `/metrics`
- Auth Service: `http://localhost:8084/health/live`, `/health/ready`, `/metrics`
- Product Swagger: `http://localhost:8081/swagger/index.html`
- Order Swagger: `http://localhost:8082/swagger/index.html`
- Consul UI: `http://localhost:8500`

## Feature notes

- Gateway and Order-to-Product calls can use healthy Consul instances with round-robin selection and configured URL fallback.
- Product calls use a request timeout, retry attempts with exponential backoff, a circuit breaker, and context cancellation.
- Product Service owns stock. Stock reservations/releases use database transactions and an idempotency ledger. Order creation attempts compensation if saving the order fails. This is a saga/compensation pattern, not a distributed ACID transaction; monitor and reconcile persistent failures.
- Metrics include request counts, error counts, latency, and database errors. Database pool sizes are configurable with `DB_MAX_OPEN_CONNS` and `DB_MAX_IDLE_CONNS`.
