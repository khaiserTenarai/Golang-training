#Day 19 Employee and Department Microservices (Go)

A beginner-friendly two-service example using Gin. Data is stored in memory, so it resets when a service restarts.

## Requirements
- Go 1.22+
- PowerShell, Command Prompt, or a terminal

## Project layout
- `employee-service`: Employee CRUD API and REST client for Department Service
- `department-service`: Department CRUD API
- `shared/dto`: API request/response DTOs, separate from entities
- `shared/utility`: correlation IDs, retry, circuit breaker
- `contract`: cross-service contract tests

## Run locally
Open two terminals from the project root.

Terminal 1:
```powershell
go run ./department-service
```
Terminal 2:
```powershell
go run ./employee-service
```

Copy `.env.example` to `.env` and adjust values if needed. The sample loads `.env` when present; environment variables already set in your shell take precedence.

## Endpoints
Department Service (default port 8082):
- `GET /health`
- `GET /departments`
- `GET /departments/:id`
- `POST /departments`
- `PUT /departments/:id`
- `DELETE /departments/:id`

Employee Service (default port 8081):
- `GET /health`
- `GET /ready` (checks Department Service)
- `GET /employees`
- `GET /employees/:id`
- `POST /employees`
- `PUT /employees/:id`
- `DELETE /employees/:id`

## Sample requests
Create a department:
```http
POST http://localhost:8082/departments
Content-Type: application/json

{"name":"IT","description":"Information Technology"}
```
Create an employee (use a department ID that exists):
```http
POST http://localhost:8081/employees
Content-Type: application/json
X-Correlation-ID: demo-123

{"name":"Anita Sharma","email":"anita@example.com","department_id":1}
```
Get employee with department details:
```http
GET http://localhost:8081/employees/1
X-Correlation-ID: demo-123
```

## Resilience and observability
- HTTP client timeout comes from `HTTP_TIMEOUT`.
- Retry count comes from `RETRY_COUNT`; retries are limited to transient failures and GET requests in the service client.
- Circuit breaker opens after `CIRCUIT_BREAKER_FAILURE_THRESHOLD` consecutive dependency failures and resets after `CIRCUIT_BREAKER_OPEN_TIMEOUT`.
- `X-Correlation-ID` is generated if absent and returned in the response. Employee Service forwards it to Department Service.
- Logs include method, path, status, duration, and correlation ID.
- `/health` reports process health; `/ready` checks the Employee Service's Department dependency.
- Error responses use JSON: `{"error":"...","correlation_id":"..."}`.

## Tests
```powershell
go test ./...
```
Contract tests use `httptest` and do not require either service to be running.

