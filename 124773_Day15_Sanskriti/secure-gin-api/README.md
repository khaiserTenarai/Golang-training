# Secure Gin Employee API

Gin + PostgreSQL (`pgxpool`) REST API demonstrating JWT authentication, role-based access control, custom middleware, validation, pagination, graceful shutdown, and `httptest` tests. Functions are grouped by layer rather than split into one file per function.

## Setup
1. Install Go 1.23+ and PostgreSQL.
2. Create the database: `CREATE DATABASE secure_employee_db;`
3. Connect to it and run `schema.sql`.
4. Copy `.env.example` to `.env`, set PostgreSQL credentials, and set `JWT_SECRET` to a random value of at least 32 characters.
5. Install dependencies and run tests:

```powershell
go mod tidy
go test ./...
go run .
```

## Hardcoded demo login credentials

The login endpoint uses these hardcoded demo accounts and generates a signed JWT with the configured `JWT_SECRET`:

```json
{"username":"admin","password":"admin123"}
```

```json
{"username":"user","password":"user123"}
```

No rows in the `users` table are needed for login. The PostgreSQL database is still required for employee CRUD operations. These hardcoded credentials are for local assignment/demo use only; use a database-backed user store and hashed passwords in a real deployment.

## Generate a JWT in SoapUI

1. Start the API with `go run .`.
2. Create a `POST` request to `http://localhost:8080/login`.
3. Select a JSON request body and send either credential object above.
4. Copy the `token` field from the response. The token is generated dynamically and expires after `JWT_EXPIRY_MINUTES` (default 60 minutes).
5. For protected requests, add the header `Authorization: Bearer <token>`.

Example successful response shape (the token value will be different each time):

```json
{
  "token": "<generated-jwt-token>",
  "token_type": "Bearer",
  "expires_in_minutes": 60,
  "username": "admin",
  "role": "admin"
}
```

## Routes
- `GET /health` public health check
- `POST /login` JSON `{"username":"admin","password":"..."}`
- `GET /employees?page=1&limit=10&sort=name&order=asc` user/admin
- `GET /employees/:id` user/admin
- `POST /employees` admin only
- `PUT /employees/:id` admin only
- `DELETE /employees/:id` admin only
- `GET /demo/next` demonstrates middleware `Next()` order
- `GET /demo/abort` demonstrates `Abort()` stopping later handlers
- `GET /ready` basic readiness endpoint (database is checked at startup)


## SoapUI example
POST `http://localhost:8080/login`, 
 eg : {"name":"Abhjnb",
"email":"abbnnm5@example.com",
"department":"IT",
"salary":45000}