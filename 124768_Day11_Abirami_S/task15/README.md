# Day 11 Mini Project — Employee REST API

Beginner-friendly PostgreSQL CRUD API using Go, Gin, and GORM.

## Requirements
- Go 1.22+
- PostgreSQL installed and running locally

## 1. Create the database
In pgAdmin or `psql`, run:

```sql
CREATE DATABASE employee_db;
```

The code defaults to host `localhost`, port `5432`, user `postgres`, password `postgres`, and database `employee_db`. If your local PostgreSQL credentials differ, set `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, and `DB_NAME` environment variables before running. `.env.example` is a reference file; the app does not automatically load `.env`.

## 2. Run the API

```bash
go mod tidy
go run .
```

The server starts at `http://localhost:8080`. GORM automatically creates the `employees` table.

## Endpoints

| Method | Endpoint | Purpose |
|---|---|---|
| GET | `/health` | Health check |
| POST | `/employees` | Create employee |
| GET | `/employees` | List employees |
| GET | `/employees/:id` | Get one employee |
| PUT | `/employees/:id` | Update employee |
| DELETE | `/employees/:id` | Delete employee |

Example JSON:

```json
{
  "name": "Rajesh",
  "email": "rajesh@example.com",
  "department": "IT",
  "salary": 75000
}
```

## Test with curl

Create:
```bash
curl -i -X POST http://localhost:8080/employees -H "Content-Type: application/json" -d '{"name":"Rajesh","email":"rajesh@example.com","department":"IT","salary":75000}'
```

List:
```bash
curl http://localhost:8080/employees
```

Get by ID:
```bash
curl http://localhost:8080/employees/1
```

Update:
```bash
curl -i -X PUT http://localhost:8080/employees/1 -H "Content-Type: application/json" -d '{"name":"Rajesh Kumar","email":"rajesh@example.com","department":"Engineering","salary":82000}'
```

Delete:
```bash
curl -i -X DELETE http://localhost:8080/employees/1
```

`name`, `email`, and `department` are required; email must be valid and unique; salary must be zero or greater. Do not use the example password in production.
