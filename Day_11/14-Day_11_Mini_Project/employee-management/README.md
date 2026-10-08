# Employee Management REST API

## Overview

This is a ready-to-run Employee Management REST API built with Go standard `net/http`, PostgreSQL and a layered architecture.

## Architecture

```text
Client
  |
  v
Logging Middleware
  |
  v
Controller
  |
  v
Service
  |
  v
Repository
  |
  v
PostgreSQL
```

## Included

- Go `net/http`
- `.env` configuration
- Config package
- Logging middleware
- PostgreSQL
- Controller layer
- Service layer
- Repository layer
- REST APIs
- JSON request and response bodies
- Path parameter
- Query parameter
- 
## Query Parameter

The collection endpoint supports a simple `sort` query parameter.

```text
GET /employees?sort=name
```

The allowed values are `name` and the default `id`.

## Path Parameter

The resource endpoint uses an employee ID path parameter.

```text
GET /employees/1
PUT /employees/1
DELETE /employees/1
```

## Setup

### 1. Create the PostgreSQL database

```sql
CREATE DATABASE employee_db;
```

### 2. Update `.env`

Change the PostgreSQL username and password if required.

### 3. Create the table

Run:

```bash
psql -U postgres -d employee_db -f schema.sql
```

### 4. Download dependencies

```bash
go mod tidy
```

### 5. Run the application

```bash
go run .
```

The server starts on:

```text
http://localhost:8080
```

## REST APIs

### Create Employee

```text
POST /employees
Content-Type: application/json
```

Request:

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "department": "IT",
  "salary": 85000
}
```

### Get All Employees

```text
GET /employees
```

### Get All Employees Sorted by Name

```text
GET /employees?sort=name
```

### Get Employee by ID

```text
GET /employees/1
```

### Update Employee

```text
PUT /employees/1
Content-Type: application/json
```

Request:

```json
{
  "name": "John Updated",
  "email": "john.updated@example.com",
  "department": "Engineering",
  "salary": 95000
}
```

### Delete Employee

```text
DELETE /employees/1
```

## Test with curl

### Create

```bash
curl -X POST http://localhost:8080/employees -H "Content-Type: application/json" -d "{\"name\":\"John Doe\",\"email\":\"john@example.com\",\"department\":\"IT\",\"salary\":85000}"
```

### Get All

```bash
curl http://localhost:8080/employees
```

### Get Sorted

```bash
curl "http://localhost:8080/employees?sort=name"
```

### Get By ID

```bash
curl http://localhost:8080/employees/1
```

### Update

```bash
curl -X PUT http://localhost:8080/employees/1 -H "Content-Type: application/json" -d "{\"name\":\"John Updated\",\"email\":\"john.updated@example.com\",\"department\":\"Engineering\",\"salary\":95000}"
```

### Delete

```bash
curl -X DELETE http://localhost:8080/employees/1
```

## Project Structure

```text
employee-management/
├── .env
├── go.mod
├── main.go
├── schema.sql
├── README.md
├── config/
│   └── config.go
├── controller/
│   └── employee_controller.go
├── logging/
│   └── logger.go
├── middleware/
│   └── logging.go
├── model/
│   └── employee.go
├── repository/
│   └── employee_repository.go
└── service/
    └── employee_service.go
```
