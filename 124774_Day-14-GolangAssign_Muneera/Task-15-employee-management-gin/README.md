# Employee Management - Gin

Basic version converted from the net/http project to Gin without adding authentication, authorization, Swagger, pagination or filtering.

## Run
1. Create database: `CREATE DATABASE employee_db;`
2. Update `.env`.
3. Run `psql -U postgres -d employee_db -f schema.sql`
4. Run `go mod tidy`
5. Run `go run .`

## APIs
GET /employees
GET /employees?sort=name
GET /employees/:id
POST /employees
PUT /employees/:id
DELETE /employees/:id

Architecture: Gin Middleware -> Controller -> Service -> Repository -> PostgreSQL.
