# Day 15 — Advanced Gin: Secure Employee API

A task-wise Go mini project demonstrating Gin middleware, execution order, `Next()`, `Abort()`, request IDs, panic recovery, CORS, JWT authentication, RBAC, request validation, pagination, file upload, and `httptest` tests.

## Requirements

- Go 1.22 or newer
- Internet access on first run so Go can download dependencies

## Run

Set a strong JWT secret before starting the API.

### PowerShell

```powershell
$env:JWT_SECRET="replace-with-a-long-random-secret"
$env:CORS_ALLOWED_ORIGINS="http://localhost:3000"
go mod tidy
go test ./...
go run .
```

### macOS / Linux

```bash
export JWT_SECRET="replace-with-a-long-random-secret"
export CORS_ALLOWED_ORIGINS="http://localhost:3000"
go mod tidy
go test ./...
go run .
```

The API listens on `http://localhost:8080` by default.

The built-in development fallback secret exists only to make local exploration easier. Do not use it in a deployed environment.

## Day 15 task map

1. Custom middleware: `internal/middleware/trace.go`
2. Execution order: `TraceOrder()` and `TestNextExecutionOrder`
3. `Next()`: `TraceOrder()`, `Timing()`, and the middleware test
4. `Abort()`: `Authenticate()`, `RequireRoles()`, and `TestAbortStopsLaterHandlers`
5. Request ID: `internal/middleware/request_id.go`
6. Recovery: `internal/middleware/recovery.go`
7. CORS: `cors.New(...)` in `internal/api/router.go`
8. JWT: `internal/middleware/auth.go`
9. Protected routes: `/api` route group
10. RBAC: `RequireRoles(...)`
11. Validation: Gin `binding` tags in `internal/model/employee.go` and request structs
12. Pagination: `GET /api/employees?page=1&limit=10`
13. File upload: `POST /api/files` example below
14. `httptest`: `internal/api/router_test.go` and `internal/middleware/middleware_test.go`

## File upload

`POST /api/files` accepts a multipart field named `file`, requires an admin or manager token, and rejects empty files or files larger than 5 MB. This demo validates metadata and returns the sanitized filename; it intentionally does not persist uploaded bytes. For production, inspect file content, restrict allowed extensions/types, enforce request-size limits at the server/proxy, and store files outside the web root or in object storage.

## Get a token

The demo token endpoint accepts a user ID and role so you can exercise RBAC locally. This is a learning-only token issuer, not real login: in a production system, authenticate credentials or an identity provider and derive roles from trusted server-side records.

```bash
curl -X POST http://localhost:8080/api/auth/token \
  -H "Content-Type: application/json" \
  -d '{"user_id":7,"role":"admin"}'
```

Copy `access_token` from the response, then use it in requests:

```bash
curl http://localhost:8080/api/employees \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## Endpoints

| Method | Path | Access |
|---|---|---|
| GET | `/health` | Public |
| POST | `/api/auth/token` | Public demo token issuer |
| GET | `/api/me` | Any valid token |
| GET | `/api/employees` | Any valid token |
| GET | `/api/employees/:id` | Any valid token |
| POST | `/api/employees` | Admin or manager |
| PUT / PATCH | `/api/employees/:id` | Admin or manager |
| DELETE | `/api/employees/:id` | Admin |
| GET | `/api/admin/overview` | Admin |
| POST | `/api/files` | Admin or manager, after adding the file-upload route above |
| GET | `/demo/next-abort` | Public middleware demo |
| GET | `/demo/panic` | Public recovery demo |

List filters: `search`, `department`, `page`, `limit`. Maximum page size is 100.

Example create body:

```json
{
  "name": "Neha Verma",
  "email": "neha@example.com",
  "department": "Engineering",
  "role": "employee"
}
```

## File upload request

After adding the route described above:

```bash
curl -X POST http://localhost:8080/api/files \
  -H "Authorization: Bearer YOUR_ADMIN_OR_MANAGER_TOKEN" \
  -F "file=@./sample.pdf"
```

## Middleware execution order

For middleware A, middleware B, and the route handler, the order is:

```text
A before
B before
handler
B after
A after
```

`c.Next()` pauses the current middleware while the remaining handlers execute, then resumes it. `c.Abort()` prevents later handlers in the chain from running; return immediately after an abort response to avoid accidentally continuing work in the current function.

## Tests

```bash
go test ./...
```

Tests cover middleware order, abort behavior, request IDs, response timing, unauthenticated access, RBAC, validation, pagination, and panic recovery.
