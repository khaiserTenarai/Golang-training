# Day 14 — Gin Framework (Piyush)

No PostgreSQL. In-memory store with 4 seed employees.

## Run
```bash
go mod tidy
go run main.go
```
Open http://localhost:8080 for the static page, or use curl for the API.

## Task Mapping

| Task | What | Where |
|------|------|-------|
| 1 | Gin application | main.go (`gin.Default()`) |
| 2 | GET routes | handler/employee.go (GetAll, GetByID) |
| 3 | POST routes | handler/employee.go (Create) |
| 4 | PUT routes | handler/employee.go (Update) |
| 5 | PATCH routes | handler/employee.go (Patch — partial update) |
| 6 | DELETE routes | handler/employee.go (Delete) |
| 7 | Route groups | main.go (`r.Group("/api")`) |
| 8 | Path parameters | handler/employee.go (`c.Param("id")`) |
| 9 | Query parameters | handler/employee.go (`c.Query`, `c.DefaultQuery`) |
| 10 | Bind JSON | handler/employee.go (`c.ShouldBindJSON`) |
| 11 | JSON responses | handler/employee.go (`c.JSON`) |
| 12 | Form parameters | handler/employee.go (CreateFromForm, `c.ShouldBind`) |
| 13 | Static files | main.go (`r.Static`, `r.StaticFile`) + static/ |
| 14 | Custom 404/405 | main.go (`r.NoRoute`, `r.NoMethod`) |
| 15 | Mini Project | Everything combined |

## Quick Test
```bash
curl http://localhost:8080/api/employees
curl http://localhost:8080/api/employees/1
curl -X POST http://localhost:8080/api/employees -H "Content-Type: application/json" -d '{"name":"Eve","email":"eve@test.com","age":26,"department":"HR","salary":52000}'
curl -X PATCH http://localhost:8080/api/employees/1 -H "Content-Type: application/json" -d '{"salary":80000}'
curl -X DELETE http://localhost:8080/api/employees/3
```
