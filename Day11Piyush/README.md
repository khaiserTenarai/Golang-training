# Day 11 — REST API with Go (Piyush)

## Structure

| Folder | Tasks | PostgreSQL? | Port |
|--------|-------|-------------|------|
| Task01_Basic_Server | 1 | No | 8080 |
| Task02_to_13_REST_API | 2–13 | No (in-memory) | 8080 |
| Task14_15_Mini_Project | 14, 15 | Yes | 8081 |

## How to Run

### Tasks 1–13 (No PostgreSQL needed)
```bash
cd Task01_Basic_Server
go run main.go

cd Task02_to_13_REST_API
go run main.go
```

### Tasks 14–15 (PostgreSQL required)
```sql
CREATE DATABASE day11db;
```
```bash
cd Task14_15_Mini_Project
go mod tidy
go run main.go
```
Edit `config/database.go` if your credentials differ from postgres/postgres.

## Task Mapping

1. **Basic HTTP server** → Task01_Basic_Server
2. **GET /employees** → Task02_to_13 (handleEmployees GET)
3. **Employee detail** → Task02_to_13 (handleEmployeeByID GET)
4. **POST** → Task02_to_13 (handleEmployees POST)
5. **PUT** → Task02_to_13 (handleEmployeeByID PUT)
6. **DELETE** → Task02_to_13 (handleEmployeeByID DELETE)
7. **Status codes** → Task02_to_13 (200, 201, 400, 404, 405, 422)
8. **JSON responses** → Task02_to_13 (APIResponse struct)
9. **JSON request bodies** → Task02_to_13 (json.NewDecoder)
10. **Path parameters** → Task02_to_13 (extractID from URL)
11. **Query parameters** → Task02_to_13 (?name, ?department, ?min_salary, ?page)
12. **Logging middleware** → Task02_to_13 (loggingMiddleware)
13. **Validation** → Task02_to_13 (validateEmployee)
14. **PostgreSQL connection** → Task14_15_Mini_Project
15. **Mini Project** → Task14_15_Mini_Project (full CRUD + PostgreSQL)

## Test Commands (Tasks 2–13, port 8080)

```bash
# List all
curl http://localhost:8080/employees

# Search
curl "http://localhost:8080/employees?name=piyush"
curl "http://localhost:8080/employees?department=Engineering"
curl "http://localhost:8080/employees?min_salary=50000"
curl "http://localhost:8080/employees?page=1&page_size=2"

# Detail
curl http://localhost:8080/employees/1

# Create
curl -X POST http://localhost:8080/employees -H "Content-Type: application/json" -d "{\"name\":\"Diana\",\"email\":\"diana@test.com\",\"age\":26,\"department\":\"HR\",\"salary\":52000}"

# Update
curl -X PUT http://localhost:8080/employees/1 -H "Content-Type: application/json" -d "{\"name\":\"Piyush Updated\",\"email\":\"p@new.com\",\"age\":26,\"department\":\"Backend\",\"salary\":70000}"

# Delete
curl -X DELETE http://localhost:8080/employees/3
```
