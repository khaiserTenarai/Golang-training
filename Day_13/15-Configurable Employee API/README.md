# Configurable Employee API — Task 15

A REST API built with Go, Gin, GORM, and PostgreSQL.

This project combines configuration management, secrets, validation,
feature flags, context propagation, logging, health checks, and graceful
shutdown.

## Features

- Environment configuration using `.env` and environment variables
- Configuration defaults and port validation
- Separate database password handling
- PostgreSQL connectivity check
- GORM automatic migration
- Layered architecture
- Employee validation in the service layer
- Gin HTTP controllers
- Database operations through a repository
- Configurable employee creation
- Application and database health endpoints
- Context propagation to database operations
- Graceful shutdown on Ctrl+C or SIGTERM
- Request logging and panic recovery

## Requirements

- Go 1.24 or later
- PostgreSQL
- A PostgreSQL database named `employee_db`

## Project structure

    employee-api/
    ├── .env
    ├── .env.example
    ├── .gitignore
    ├── go.mod
    ├── README.md
    ├── main.go
    ├── config/
    ├── database/
    ├── model/
    ├── validation/
    ├── repository/
    ├── service/
    ├── controller/
    └── server/

## Setup

### 1. Create your environment file

In PowerShell, run:

    Copy-Item .env.example .env

Edit `.env` and set `DB_PASSWORD` to your local PostgreSQL password.

Never commit `.env` or real credentials to source control.

### 2. Create the database

Create the database in PostgreSQL:

    CREATE DATABASE employee_db;

Update the database settings in `.env` if your host, port, username,
or database name is different.

### 3. Install dependencies

Run these commands from the folder containing `go.mod`:

    go mod tidy
    go fmt ./...
    go build ./...

### 4. Start the application

    go run .

The default server address is:

    http://localhost:8080

## Configuration

| Variable | Default | Purpose |
|---|---|---|
| APP_ENV | development | Application environment |
| APP_PORT | 8080 | HTTP port |
| DB_HOST | localhost | PostgreSQL host |
| DB_PORT | 5432 | PostgreSQL port |
| DB_USER | postgres | PostgreSQL username |
| DB_PASSWORD | Required | PostgreSQL password |
| DB_NAME | employee_db | Database name |
| DB_SSLMODE | disable | PostgreSQL SSL mode |
| ENABLE_EMPLOYEE_CREATE | true | Enable or disable employee creation |

In production, use an appropriate PostgreSQL SSL mode and provide
secrets through your deployment environment or secret manager.

## API endpoints

| Method | Endpoint | Description |
|---|---|---|
| GET | /health | Application health |
| GET | /health/db | Database connectivity |
| POST | /employees | Create an employee if enabled |
| GET | /employees/:id | Get an employee by ID |

## Test the application health

    Invoke-RestMethod http://localhost:8080/health

Expected response:

    {
      "status": "ok"
    }

## Test database health

    Invoke-RestMethod http://localhost:8080/health/db

## Create an employee

Make sure `ENABLE_EMPLOYEE_CREATE=true` and restart the application.

In PowerShell:

    Invoke-RestMethod `
      -Method Post `
      -Uri "http://localhost:8080/employees" `
      -ContentType "application/json" `
      -Body '{"name":"Ganesh","email":"ganesh@example.com"}'

A successful request returns HTTP 201 Created.

## Retrieve an employee

Use the ID returned from the create request:

    Invoke-RestMethod http://localhost:8080/employees/1

Replace `1` with the actual employee ID.

## Feature flag

### Enable employee creation

Set the following in `.env`:

    ENABLE_EMPLOYEE_CREATE=true

Restart the application.

`POST /employees` is registered and employee creation is available.

### Disable employee creation

Set:

    ENABLE_EMPLOYEE_CREATE=false

Restart the application.

`POST /employees` returns HTTP 404 because the route is not registered.

The health endpoints and `GET /employees/:id` remain available.

The feature flag is read at startup. Changing `.env` requires a restart.

## Architecture

The request flow is:

    HTTP request
        |
        v
    Controller
        |
        v
    Service
        |
        v
    Validation
        |
        v
    Repository
        |
        v
    GORM
        |
        v
    PostgreSQL

The controller handles HTTP requests and responses.
The service owns business logic and calls validation.
The repository performs database operations.

The request context is passed to GORM so cancellation and deadlines
can propagate to database calls.

## Startup sequence

1. Load and validate configuration.
2. Load database secrets.
3. Connect to PostgreSQL.
4. Verify database connectivity.
5. Run database migration.
6. Initialize repository, service, and controller.
7. Register routes according to the feature flag.
8. Start the HTTP server.

## Shutdown sequence

When Ctrl+C or SIGTERM is received:

1. The application receives the shutdown signal.
2. The HTTP server starts graceful shutdown.
3. Active HTTP requests are given up to 10 seconds to finish.
4. The HTTP server stops.
5. The database connection pool is closed.

## Troubleshooting

### go.mod not found

Run Go commands from the project root, where `go.mod` exists.

If this is a new project and no `go.mod` exists, run:

    go mod init employee-api

Then run:

    go mod tidy

Keep this project separate from sibling assignment directories whose names
contain spaces.

### Database connection failed

Check that:

- PostgreSQL is running.
- The database exists.
- DB_HOST and DB_PORT are correct.
- DB_USER and DB_PASSWORD are correct.
- DB_NAME matches the database you created.
- DB_SSLMODE matches your PostgreSQL setup.

### POST /employees returns 404

Check that `ENABLE_EMPLOYEE_CREATE=true` is in `.env`, then restart
the application.

### Duplicate email

The employee email column has a unique database index. Use a different
email address for another successful creation request.

## Security notes

- Never commit `.env` or real passwords.
- Use a secret manager or deployment environment for production secrets.
- Configure PostgreSQL SSL appropriately in production.
- Add authentication, authorization, rate limiting, and production-grade
  error handling before exposing this learning project publicly.