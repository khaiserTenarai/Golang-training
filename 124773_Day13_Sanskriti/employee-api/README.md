# Employee API

Configurable Go service with environment configuration, logging, context propagation, and graceful shutdown.

## Getting Started

1. Copy `.env.example` to `.env`:
   ```bash
   cp .env.example .env
   ```

2. Run the application:
   ```bash
   go run cmd/server/main.go
   ```

3. Run unit tests:
   ```bash
   go test -v ./tests/...
   ```

## Endpoints

- `POST /api/v1/employees` - Register a new employee
- `GET /api/v1/employees/:id` - Fetch an employee by ID
