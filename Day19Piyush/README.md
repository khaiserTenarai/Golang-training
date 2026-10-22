# Day 19 — Go Microservices

**Author:** Piyush  
**Date:** Day 19 of Golang Training  
**Topic:** Microservices with Resilience and Observability

---

## Overview

This project implements two interconnected Go microservices — **Employee Service** and **Department Service** — demonstrating production-grade patterns like service-to-service REST communication, DTO/Entity separation, retry with exponential backoff, circuit breaker, correlation ID propagation, externalized configuration, health checks, and contract testing.

No PostgreSQL or external dependencies required. All data is stored in-memory.

## Project Structure

```
Day19Piyush/
├── go.mod                          # Go module definition
├── .gitignore                      # Git ignore rules
├── README.md                       # This file
├── decisions.md                    # Architecture decisions and reasoning
├── config/
│   └── config.go                   # Task 7: Externalized service URLs
├── entity/
│   └── entity.go                   # Task 6: Internal domain entities
├── dto/
│   └── dto.go                      # Tasks 5, 6: Data Transfer Objects
├── middleware/
│   └── correlation.go              # Task 12: Correlation ID middleware
├── resilience/
│   ├── retry.go                    # Task 11: Retry with exponential backoff
│   └── circuit_breaker.go          # Task 13: Circuit breaker (CLOSED/OPEN/HALF-OPEN)
├── httpclient/
│   └── client.go                   # Tasks 4, 9, 10: HTTP client with timeout
├── employee-service/
│   └── main.go                     # Tasks 1, 3, 4, 7, 8: Employee microservice
├── department-service/
│   └── main.go                     # Tasks 2, 3, 8: Department microservice
└── tests/
    └── contract_test.go            # Task 14: Contract tests (7 tests)
```

## Quick Start

```bash
# Terminal 1 — Start Department Service
go run ./department-service

# Terminal 2 — Start Employee Service
go run ./employee-service

# Terminal 3 — Run Tests
go test -v ./tests/...
```

## Endpoints

### Employee Service (Port 8081)
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /employees | List all (enriched with department) |
| GET | /employees/{id} | Get by ID |
| POST | /employees | Create employee |
| PUT | /employees/{id} | Update employee |
| DELETE | /employees/{id} | Delete employee |
| GET | /health | Health check + circuit breaker state |

### Department Service (Port 8082)
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /departments | List all |
| GET | /departments/{id} | Get by ID |
| POST | /departments | Create department |
| DELETE | /departments/{id} | Delete department |
| GET | /health | Health check |

## Task Coverage

| # | Task | Implementation |
|---|------|---------------|
| 1 | Employee microservice | `employee-service/main.go` |
| 2 | Department microservice | `department-service/main.go` |
| 3 | CRUD APIs | Both services — full GET/POST/PUT/DELETE |
| 4 | Service-to-service REST | `httpclient/client.go` — employee calls department |
| 5 | Create DTOs | `dto/dto.go` — request/response DTOs |
| 6 | Separate DTOs from entities | `entity/entity.go` vs `dto/dto.go` |
| 7 | Externalize service URLs | `config/config.go` — env var based |
| 8 | Health checks | `/health` endpoint on both services |
| 9 | Service error handling | `httpclient/client.go` + handler error responses |
| 10 | HTTP client timeout | `httpclient/client.go` — 5s configurable timeout |
| 11 | Retry | `resilience/retry.go` — exponential backoff (3 attempts) |
| 12 | Correlation ID | `middleware/correlation.go` — X-Correlation-ID header |
| 13 | Circuit breaker | `resilience/circuit_breaker.go` — state machine |
| 14 | Contract tests | `tests/contract_test.go` — 7 tests |
| 15 | Mini Project | All above combined into working system |

## Test Commands

```bash
# List employees (enriched with department)
curl http://localhost:8081/employees

# Create employee
curl -X POST http://localhost:8081/employees \
  -H "Content-Type: application/json" \
  -d '{"name":"Eve","email":"eve@test.com","age":26,"department_id":2,"salary":55000}'

# Test correlation ID propagation
curl -H "X-Correlation-ID: my-trace-123" http://localhost:8081/employees/1

# Health check (shows circuit breaker state)
curl http://localhost:8081/health

# Run contract tests
go test -v ./tests/...
```
