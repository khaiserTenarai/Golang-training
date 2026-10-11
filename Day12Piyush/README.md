# Day 12 — Production REST API Architecture (Piyush)

No PostgreSQL needed. In-memory store with 5 seed employees.

## Run
```bash
go mod tidy
go run main.go
```
Server starts on http://localhost:8080

## Architecture (Tasks 1–4)
```
main.go                          → Wiring + DI (Task 1, 4)
├── routes/routes.go             → Route definitions + middleware chain
├── middleware/
│   ├── request_id.go            → Task 5: X-Request-ID header
│   ├── logging.go               → Task 6: Request logging
│   ├── auth.go                  → Task 7: Authentication (Bearer token)
│   │                              Task 8: Authorization (admin vs user)
│   └── recovery.go              → Panic recovery
├── handler/employee_handler.go  → Task 2: HTTP handler layer
├── service/employee_service.go  → Task 2, 3: Business logic layer
├── repository/employee_repo.go  → Task 3: Data access layer (interface)
├── validator/employee_validator.go → Task 9: Validation layer
└── models/employee.go           → Shared data structures
```

## Task Mapping

| Task | What | Where |
|------|------|-------|
| 1 | Layered architecture | Entire project structure |
| 2 | Separate Handler & Service | handler/ + service/ |
| 3 | Separate Service & Repository | service/ + repository/ |
| 4 | Dependency injection | Interfaces in service + repository |
| 5 | Request-ID middleware | middleware/request_id.go |
| 6 | Logging middleware | middleware/logging.go |
| 7 | Authentication | middleware/auth.go (AuthMiddleware) |
| 8 | Authorization | middleware/auth.go (AdminWriteOnly) |
| 9 | Validation layer | validator/employee_validator.go |
| 10 | Pagination | ?page=1&page_size=2 |
| 11 | Filtering | ?name=, ?department=, ?min_salary=, ?max_salary= |
| 12 | Sorting | ?sort_by=salary&sort_order=desc |
| 13 | /health endpoint | GET /health |
| 14 | /ready endpoint | GET /ready |
| 15 | Mini Project | Everything combined |

## Auth Tokens
| Token | Role | Access |
|-------|------|--------|
| `admin-token-123` | admin | Read + Write (full CRUD) |
| `user-token-456` | user | Read only |
| `user-token-789` | user | Read only |

## Quick Test
```bash
# Health (public)
curl http://localhost:8080/health

# List (user token)
curl -H "Authorization: Bearer user-token-456" http://localhost:8080/api/employees

# Create (admin only)
curl -X POST http://localhost:8080/api/employees \
  -H "Authorization: Bearer admin-token-123" \
  -H "Content-Type: application/json" \
  -d '{"name":"Eve","email":"eve@test.com","age":26,"department":"Finance","salary":52000,"role":"developer"}'
```
