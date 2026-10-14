# Day 14 Mini Project - Gin Employee CRUD REST API

A simple Employee CRUD REST API using Gin and an in-memory slice.

## Endpoints

- GET `/employees` - View all employees
- GET `/employees/:id` - View one employee
- POST `/employees` - Create employee
- PUT `/employees/:id` - Replace employee
- PATCH `/employees/:id` - Partially update employee
- DELETE `/employees/:id` - Delete employee

## Run

```bash
go mod tidy
go run .
```

Server runs on `http://localhost:8080`.

## Example JSON

### POST /employees

```json
{
  "name": "Abirami",
  "department": "IT",
  "salary": 40000
}
```

### PUT /employees/1

```json
{
  "name": "Abirami S",
  "department": "Engineering",
  "salary": 45000
}
```

### PATCH /employees/1

```json
{
  "salary": 50000
}
```
