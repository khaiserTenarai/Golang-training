# Production-style Layered Employee API

## Requirements
- Go 1.22+
- No third-party dependencies

## Run
```bash
go run ./cmd/api
```

The server listens on `:8080` by default. Set `PORT` to change the port and `JWT_SECRET` to set the signing secret.

## Demo authentication
The API uses HMAC-signed bearer tokens. For local development, request a token:

```bash
curl -X POST http://localhost:8080/auth/token \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}'
```

The demo credentials are intentionally simple and must be replaced with a real identity provider or securely stored credentials before production use.

Use the returned token:
```bash
curl http://localhost:8080/api/employees \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## Endpoints
- `GET /health`
- `GET /ready`
- `POST /auth/token`
- `POST /api/employees`
- `GET /api/employees`
- `GET /api/employees/{id}`
- `PUT /api/employees/{id}`
- `DELETE /api/employees/{id}`

Employee list query parameters:
- `page` (default 1)
- `limit` (default 10, maximum 100)
- `department` (case-insensitive exact match)
- `search` (matches name or email)
- `sort_by` (`name`, `email`, `department`, `created_at`, `id`)
- `order` (`asc` or `desc`)

All `/api/*` routes require a valid bearer token. Creating, updating, and deleting employees require the `admin` role. The demo token endpoint issues an `admin` token for the demo account and a `user` token for the user account.

## Task-to-code map
1. Layered architecture: `internal/handler`, `internal/service`, `internal/repository`, `internal/domain`
2. Handler and Service separation: `internal/handler/employee_handler.go`, `internal/service/employee_service.go`
3. Service and Repository separation: `internal/service/employee_service.go`, `internal/repository/employee_repository.go`
4. Dependency injection: `cmd/api/main.go`
5. Request-ID middleware: `internal/middleware/request_id.go`
6. Logging middleware: `internal/middleware/logging.go`
7. Authentication: `internal/auth/token.go`, `internal/handler/auth_handler.go`, `internal/middleware/auth.go`
8. Authorization: `internal/middleware/authorization.go`
9. Validation layer: `internal/validation/employee.go`
10. Pagination: `internal/service/employee_service.go`
11. Filtering: `internal/service/employee_service.go`
12. Sorting: `internal/service/employee_service.go`
13. Health endpoint: `internal/handler/system_handler.go`
14. Readiness endpoint: `internal/handler/system_handler.go`
15. Mini project composition: `cmd/api/main.go`, `internal/router/router.go`

The repository is in-memory to keep the project runnable without external services. Replace it with a database-backed implementation while keeping the same repository interface for persistent storage.
