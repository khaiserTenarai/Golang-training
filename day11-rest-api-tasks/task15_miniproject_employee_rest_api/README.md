# Employee REST API (Go + PostgreSQL)

Complete CRUD API with validation, filtering, sorting, pagination, logging middleware, panic recovery, and graceful shutdown.

## Run

This project has several files, so run it with `go run .` (not `go run main.go`).

1. Create the database (once):

```powershell
psql -U postgres -c "CREATE DATABASE employeedb;"
```

2. If your PostgreSQL password is not `password`, set the connection string:

```powershell
$env:DATABASE_URL = "postgres://postgres:YOUR_PASSWORD@localhost:5432/employeedb?sslmode=disable"
```

3. Start the server. The `employees` table is created automatically on startup (`init.sql` is the same schema for reference):

```powershell
go run .
```

Optional: set `$env:PORT = "9090"` to change the port.

## Endpoints

| Method | Path | Success | Errors |
|--------|------|---------|--------|
| GET | `/health` | 200 | 503 if database is down |
| GET | `/employee` | 200 | 400 bad query params |
| POST | `/employee` | 201 + `Location` header | 400, 409 duplicate email, 415, 422 validation |
| GET | `/employee/{id}` | 200 | 400, 404 |
| PUT | `/employee/{id}` | 200 | 400, 404, 409, 415, 422 |
| DELETE | `/employee/{id}` | 204 | 400, 404 |

List query parameters: `page`, `limit` (1-100, default 10), `department`, `search` (name or email), `sort` (`id`, `name`, `salary`, `department`, `created_at`), `order` (`asc` or `desc`).

Allowed departments: Engineering, Finance, HR, Marketing, Sales.

## Try it (PowerShell)

```powershell
Invoke-RestMethod -Method Post -Uri http://localhost:8080/employee -ContentType "application/json" -Body '{"name":"Anita Rao","email":"anita@example.com","department":"Engineering","salary":85000}'

Invoke-RestMethod "http://localhost:8080/employee"

Invoke-RestMethod "http://localhost:8080/employee?department=Engineering&sort=salary&order=desc&limit=5"

Invoke-RestMethod "http://localhost:8080/employee/1"

Invoke-RestMethod -Method Put -Uri http://localhost:8080/employee/1 -ContentType "application/json" -Body '{"name":"Anita Rao","email":"anita@example.com","department":"Finance","salary":95000}'

Invoke-RestMethod -Method Delete -Uri http://localhost:8080/employee/1
```

`Invoke-RestMethod` throws on 4xx/5xx responses. To see the status code and headers, use `curl.exe` (with the `.exe`, otherwise PowerShell uses its own `curl` alias):

```powershell
curl.exe -i http://localhost:8080/employee/999
```

## Test

Unit tests use an in-memory fake repository and need no database:

```powershell
go test ./...
```

The PostgreSQL integration test runs only when `TEST_DATABASE_URL` is set. It drops and recreates the `employees` table, so point it at a separate test database, never your real data:

```powershell
psql -U postgres -c "CREATE DATABASE employeedb_test;"
$env:TEST_DATABASE_URL = "postgres://postgres:YOUR_PASSWORD@localhost:5432/employeedb_test?sslmode=disable"
go test ./...
```

## Layout

- `main.go`: configuration, database connection, server start, graceful shutdown
- `model.go`: types, repository interface, input validation
- `repository.go`: PostgreSQL implementation
- `handlers.go`: HTTP handlers, JSON helpers, error mapping, middleware, router
- `main_test.go`: handler tests and the PostgreSQL integration test
