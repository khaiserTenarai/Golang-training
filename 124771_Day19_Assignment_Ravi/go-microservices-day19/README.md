# Day 19 — Go Microservices (Employee + Department)

Beginner-friendly, layered Go microservices using Gin, GORM and PostgreSQL. Each service is a separate Go module and process. Employee Service calls Department Service over REST.

## Features / assignment mapping

1. Employee microservice — `employee-service/`
2. Department microservice — `department-service/`
3. CRUD APIs — `handler/`, `service/`, `repository/`
4. REST service-to-service call — `employee-service/client/department_client.go`
5–6. DTOs separate from database entities — each service's `dto/` and `model/`
7. Externalized URLs and settings — each service `.env.example` and `config/config.go`
8. Health checks — `/health` and `/ready`
9. Consistent service errors — `internal/httpx/` and handlers
10. HTTP client timeout — `client/department_client.go`
11. Bounded retry for transient Department Service failures — same client
12. Correlation ID propagated between services — `middleware/correlation_id.go` and client
13. Circuit breaker — `resilience/circuit_breaker.go` (simple in-memory breaker)
14. Contract tests — `tests/department_contract_test.go` in Employee Service

## Requirements

- Go 1.22+ (Go 1.23+ also fine)
- PostgreSQL 14+
- PowerShell/terminal and Postman or curl

## 1. Create database

Create one PostgreSQL database named `microservices_db`:

```sql
CREATE DATABASE microservices_db;
```

The two services use the same database for this beginner project but own separate tables (`employees` and `departments`). In a production system, each microservice would usually own its data store and other services would access it only through APIs.

## 2. Configure environment


Employee service defaults:
- `PORT=8081`
- `DEPARTMENT_SERVICE_URL=http://localhost:8082`

Department service defaults:
- `PORT=8082`

## 3. Start Department Service

Terminal 1:

```bash
cd department-service
go mod tidy
go run .
```

## 4. Start Employee Service

Terminal 2:

```bash
cd employee-service
go mod tidy
go run .
```

Both services run independently. GORM AutoMigrate creates/updates the required tables on startup. Seed one department through the API before creating an employee.

## API reference

Base URLs:
- Employee: `http://localhost:8081`
- Department: `http://localhost:8082`

### Department APIs

| Method | URL | Purpose |
|---|---|---|
| GET | `/health` | Process health |
| GET | `/ready` | Database readiness |
| POST | `/departments` | Create department |
| GET | `/departments` | List departments |
| GET | `/departments/:id` | Get one department |
| PUT | `/departments/:id` | Update department |
| DELETE | `/departments/:id` | Delete department |

Create department — `POST http://localhost:8082/departments`, JSON:

