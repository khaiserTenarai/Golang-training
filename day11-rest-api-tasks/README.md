# Day 11: REST API with Go

| Task | Topic | Needs PostgreSQL |
|------|-------|------------------|
| task01 | Basic HTTP server | no |
| task02 | GET /employee | no |
| task03 | Employee detail endpoint | no |
| task04 | POST | no |
| task05 | PUT | no |
| task06 | DELETE | no |
| task07 | Correct HTTP status codes | no |
| task08 | JSON responses | no |
| task09 | JSON request bodies | no |
| task10 | Path parameters | no |
| task11 | Query parameters | no |
| task12 | Logging middleware | no |
| task13 | Data validation | no |
| task14 | PostgreSQL connection | yes |
| task15 | Mini project: PostgreSQL-backed CRUD API | yes |

## Run a task

Tasks 1 to 14 are a single `main.go`:

```powershell
cd task07_http_status_codes
go run main.go
```

Task 15 has several files:

```powershell
cd task15_miniproject_employee_rest_api
go run .
```

Every server listens on port 8080, so stop one (Ctrl+C) before starting another.

## PostgreSQL (tasks 14 and 15)

```powershell
psql -U postgres -c "CREATE DATABASE employeedb;"
$env:DATABASE_URL = "postgres://postgres:YOUR_PASSWORD@localhost:5432/employeedb?sslmode=disable"
```

The default connection string is `postgres://postgres:password@localhost:5432/employeedb?sslmode=disable`. Both tasks create the `employees` table automatically on startup.

Both use the standard `database/sql` package with the `lib/pq` driver. The `go.sum` file is included, so `go run` works without extra steps (`go mod tidy` is also fine).
