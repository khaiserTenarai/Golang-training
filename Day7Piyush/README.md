# Day 7 - Go + PostgreSQL Tasks (Piyush)

## Prerequisites
- **Go** 1.21+ installed
- **PostgreSQL** running on localhost:5432
- Database: `day7db` (create it manually: `CREATE DATABASE day7db;`)
- Default credentials: user=`postgres`, password=`postgres`

## Setup (Run Once)

### 1. Create the database
```sql
CREATE DATABASE day7db;
```

### 2. For each task, run:
```bash
cd TaskXX_Folder
go mod tidy
go run main.go
```

> `go mod tidy` downloads all required dependencies (lib/pq, gorilla/mux, bcrypt, etc.)

## DB Credentials
All tasks use these defaults (edit `config/database.go` in any task to change):
- Host: localhost
- Port: 5432
- User: postgres
- Password: postgres
- DB Name: day7db

---

## Task List

| # | Task | Type | Port |
|---|------|------|------|
| 1 | Employee CRUD | Console | - |
| 2 | Department Management (Foreign Keys) | Console | - |
| 3 | Employee Search (Pagination + Sorting) | Console | - |
| 4 | Salary Management (Transactions) | Console | - |
| 5 | Product Inventory (Stock Mgmt) | Console | - |
| 6 | Customer Management (Search + Pagination) | Console | - |
| 7 | Bank Account System (Deposit/Withdraw) | Console | - |
| 8 | Money Transfer (Commit/Rollback) | Console | - |
| 9 | E-Commerce Order System (JOINs) | Console | - |
| 10 | Order & Inventory Transaction (Rollback) | Console | - |
| 11 | Employee REST API | REST API | 8080 |
| 12 | Repository-Service Architecture (DI) | REST API | 8081 |
| 13 | Authentication System (bcrypt + Roles) | REST API | 8082 |
| 14 | Attendance & Leave Management | REST API | 8083 |
| 15 | Employee Management Capstone | REST API | 8084 |

## Architecture (Each Task)

### Console Tasks (1-10)
```
config/      → Database connection + table creation
models/      → Data structures
repository/  → SQL queries (CRUD operations)
controller/  → Business logic (orchestrates view + repo)
view/        → Console UI (menus, input, output)
main.go      → Entry point + wiring
```

### REST API Tasks (11-15)
```
config/      → Database connection + table creation
models/      → Data structures + JSON tags
repository/  → SQL queries (interface + implementation)
service/     → Business logic (interface + implementation)
controller/  → HTTP handlers (JSON request/response)
middleware/  → Auth, logging (Tasks 13-15)
routes/      → Route definitions
main.go      → Entry point + dependency injection
```

## Testing REST APIs (Tasks 11-15)
Use curl or Postman:

```bash
# Register (Task 13/15)
curl -X POST http://localhost:8082/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"piyush","email":"piyush@test.com","password":"secret123"}'

# Login
curl -X POST http://localhost:8082/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"piyush","password":"secret123"}'

# Use token for protected routes
curl http://localhost:8082/api/auth/profile \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```
