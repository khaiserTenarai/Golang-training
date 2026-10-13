# Day 14 — Gin Employee CRUD REST API

## Requirements
- Go 1.22+
- Internet access the first time you run `go mod tidy` to download Gin and its dependencies

## Run
```bash
go mod tidy
go run .
```

The API listens on `:8080` by default. Set `PORT` to change the port.

## Endpoints

| Method | Route | Purpose |
|---|---|---|
| GET | `/health` | Health check |
| GET | `/api/employees` | List employees |
| GET | `/api/employees/:id` | Get one employee |
| POST | `/api/employees` | Create employee from JSON |
| POST | `/api/employees/form` | Create employee from form fields |
| PUT | `/api/employees/:id` | Replace employee fields |
| PATCH | `/api/employees/:id` | Update selected employee fields |
| DELETE | `/api/employees/:id` | Delete employee |
| GET | `/assets/*filepath` | Serve files from `./static` |

## Examples

Create:
```bash
curl -X POST http://localhost:8080/api/employees -H "Content-Type: application/json" -d '{"name":"Aarav Sharma","email":"aarav@example.com","department":"Engineering"}'
```

List with query parameters:
```bash
curl "http://localhost:8080/api/employees?department=Engineering&search=Aarav"
```

Get one:
```bash
curl http://localhost:8080/api/employees/1
```

Replace:
```bash
curl -X PUT http://localhost:8080/api/employees/1 -H "Content-Type: application/json" -d '{"name":"Aarav Sharma","email":"aarav@example.com","department":"Platform"}'
```

Patch:
```bash
curl -X PATCH http://localhost:8080/api/employees/1 -H "Content-Type: application/json" -d '{"department":"SRE"}'
```

Create using form fields:
```bash
curl -X POST http://localhost:8080/api/employees/form -F "name=Meera Singh" -F "email=meera@example.com" -F "department=HR"
```

## Task-to-code map
1. Gin app: `main.go`
2. GET routes: `GET /api/employees`, `GET /api/employees/:id`
3. POST routes: `POST /api/employees`, `POST /api/employees/form`
4. PUT route: `PUT /api/employees/:id`
5. PATCH route: `PATCH /api/employees/:id`
6. DELETE route: `DELETE /api/employees/:id`
7. Route groups: `/api`
8. Path parameters: `:id`
9. Query parameters: `department`, `search`
10. Gin JSON binding: `ShouldBindJSON`
11. JSON responses: `c.JSON`
12. Form parameters: `ShouldBind`
13. Static files: `./static` mounted at `/assets`
14. Custom 404/405: `NoRoute`, `NoMethod`, and `HandleMethodNotAllowed`
15. CRUD mini project: all employee routes

The in-memory store is suitable for practice. Records reset whenever the server restarts.
