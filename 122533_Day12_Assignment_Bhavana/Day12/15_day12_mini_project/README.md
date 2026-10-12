# Day 12 Mini Project - Layered Employee API

```
handler/      HTTP only (read request, write response, pick status code)
service/      business rules + validation
repository/   all the SQL (PostgreSQL)
middleware/   request ID, logging, authentication, authorization
validation/   the validation rules
model/        the Employee struct and the list filter
```

## Run

```bash
go mod tidy
go run .
```

You'll be asked for the database details, the port, and two tokens
(admin = read/write, user = read only).

## Try it

```bash
# health and ready need no token
curl localhost:8080/health
curl localhost:8080/ready

# create (admin only)
curl -X POST localhost:8080/employee -H "Authorization: Bearer <admin token>" \
  -d '{"name":"Ravi Kumar","email":"ravi@example.com","department":"IT","salary":45000}'

# list with filtering, sorting and pagination (any valid token)
curl "localhost:8080/employee?department=IT&min_salary=30000&sort_by=salary&order=desc&page=1&limit=5" \
  -H "Authorization: Bearer <user token>"

# a user token trying to delete gets 403
curl -X DELETE localhost:8080/employee/1 -H "Authorization: Bearer <user token>"
```
