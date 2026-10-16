# Day 18 — Microservices Architecture (Piyush)

No PostgreSQL. All in-memory. 4 runnable services + event bus demo + architecture docs.

## How to Run (4 terminals)

```bash
# Terminal 1
cd auth-service && go run main.go          # Port 8083

# Terminal 2
cd department-service && go run main.go    # Port 8082

# Terminal 3
cd employee-service && go run main.go      # Port 8081

# Terminal 4
cd api-gateway && go run main.go           # Port 8080 (entry point)
```

### Event Bus + Saga Demo (standalone)
```bash
cd event-bus && go run main.go
```

## Test via Gateway (port 8080)

```bash
# Login
curl -X POST http://localhost:8080/auth/login -H "Content-Type: application/json" -d "{\"username\":\"admin\",\"password\":\"admin123\"}"

# Use token from response:
TOKEN=<paste_token>

# List employees (calls employee-service which calls department-service)
curl -H "Authorization: Bearer $TOKEN" http://localhost:8080/employees

# List departments
curl -H "Authorization: Bearer $TOKEN" http://localhost:8080/departments

# Create employee (admin only)
curl -X POST http://localhost:8080/employees -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" -d "{\"name\":\"Eve\",\"email\":\"eve@test.com\",\"age\":26,\"department_id\":2,\"salary\":55000}"

# Service registry
curl http://localhost:8080/services
```

## Task Mapping

| Task | What | Where |
|------|------|-------|
| 1 | Analyze monolith | docs/architecture_design.txt |
| 2 | Microservice boundaries | docs/architecture_design.txt |
| 3 | Bounded contexts | docs/architecture_design.txt |
| 4 | Employee Service | employee-service/main.go (:8081) |
| 5 | Department Service | department-service/main.go (:8082) |
| 6 | Auth Service | auth-service/main.go (:8083) |
| 7 | Database-per-service | Each service has own data store |
| 8 | Sync REST communication | employee-service calls dept-service; gateway proxies |
| 9 | Async communication | event-bus/main.go (pub/sub event bus) |
| 10 | API Gateway | api-gateway/main.go (:8080) |
| 11 | Service discovery | api-gateway (registry + health checks) |
| 12 | Distributed transactions | event-bus/main.go (saga with failure rollback) |
| 13 | Saga workflow | event-bus/main.go (onboarding saga) |
| 14 | Observability | Logging in all services + /health endpoints |
| 15 | Mini Project | docs/architecture_design.txt (full technical design) |

## Architecture

```
Client → API Gateway (:8080)
              ├── /auth/*        → Auth Service (:8083)
              ├── /employees/*   → Employee Service (:8081)
              │                       └── calls Dept Service for enrichment
              └── /departments/* → Department Service (:8082)
              
         Event Bus (async): employee.created → [Dept, Email, IT services]
```
