# Day 15 — Advanced Gin (Piyush)

No PostgreSQL. In-memory store. JWT auth with RBAC.

## Run
```bash
go mod tidy
go run main.go
```

## Run Tests (Task 14)
```bash
go test -v ./handler/...
```

## Task Mapping

| Task | What | File |
|------|------|------|
| 1 | Custom middleware | middleware/custom.go (TimingMiddleware) |
| 2 | Middleware order | middleware/custom.go (First/Second) + GET /order |
| 3 | Next() | middleware/custom.go (code before/after c.Next) |
| 4 | Abort() | middleware/custom.go (ApiKeyMiddleware) |
| 5 | Request ID | middleware/request_id.go |
| 6 | Recovery | middleware/recovery.go |
| 7 | CORS | middleware/cors.go |
| 8 | JWT auth | middleware/auth.go (GenerateJWT, JWTAuth) |
| 9 | Protect routes | main.go (api.Use JWTAuth) |
| 10 | RBAC | middleware/auth.go (RequireRole) |
| 11 | Validation | handler/employee.go (Gin binding tags) |
| 12 | Pagination | handler/employee.go (GetAll) |
| 13 | File upload | handler/employee.go (UploadFile) |
| 14 | httptest | handler/employee_test.go |
| 15 | Mini Project | Everything combined |

## Seeded Users
| Username | Password | Role | Access |
|----------|----------|------|--------|
| admin | admin123 | admin | Full CRUD + upload |
| manager | manager123 | manager | CRUD (no upload) |
| piyush | piyush123 | user | Read only |
