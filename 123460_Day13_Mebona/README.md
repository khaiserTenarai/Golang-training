# Day 13 - Configurable Employee API (standard library only)

No Gin, GORM or GraphQL. Only `net/http`, `log/slog` and friends - `go.mod` has no dependencies.
Needs Go 1.22 or newer (method + `{id}` patterns in `http.ServeMux`).

    go test ./...
    go run .

## Layout

    .env                  base settings (no secrets)
    .env.test             overrides for APP_ENV=test
    .env.production       overrides for APP_ENV=production
    .env.secrets          secrets only, git-ignored  (.env.secrets.example = template)
    config/               struct, defaults per environment, loading, validation, Secret type
    model/                Employee, input + validation, errors
    repository/           storage interface + in-memory implementation (context everywhere)
    service/              business rules, feature flag, timeouts
    controller/           HTTP handlers, health/ready, config view
    middleware/           request id, logging, panic recovery, API-key check
    main.go               startup sequence, signal handling, shutdown sequence

## Configuration priority (highest first)

1. real environment variables   `$env:SERVER_PORT="9000"`
2. `.env.secrets`
3. `.env.<APP_ENV>`
4. `.env`
5. built-in defaults of the environment (`config/profiles.go`)

Switch environment (PowerShell):

    $env:APP_ENV="test";        go run .
    $env:APP_ENV="production";  $env:API_KEY="a-long-random-key-123456"; go run .
    Remove-Item Env:APP_ENV

Production refuses to start without a real API_KEY (16+ characters), with LOG_LEVEL=debug, or with seed data.
All configuration mistakes are reported together, then the process exits.

## Try it (PowerShell)

    Invoke-RestMethod http://localhost:8080/health
    Invoke-RestMethod http://localhost:8080/config            # secrets show as [REDACTED]
    Invoke-RestMethod "http://localhost:8080/employees?page=1&pageSize=2"
    Invoke-RestMethod "http://localhost:8080/employees?search=priya"

    $h = @{ "X-API-Key" = "my-local-dev-api-key" }            # value from .env.secrets
    $body = '{"name":"Riya","email":"riya@example.com","department":"IT","salary":82000}'
    Invoke-RestMethod -Method Post -Uri http://localhost:8080/employees -Headers $h -ContentType "application/json" -Body $body
    Invoke-RestMethod -Method Delete -Uri http://localhost:8080/employees/1 -Headers $h

Without the header you get `401`.

## Feature flags

| Variable | Effect when false |
|---|---|
| `FEATURE_SEARCH` | `?search=` returns 404 "search: feature disabled" |
| `FEATURE_EXPOSE_SALARY` | `salary` is removed from every response |
| `FEATURE_SEED_DATA` | no sample employees at startup |

    $env:FEATURE_EXPOSE_SALARY="false"; go run .

## Graceful shutdown demo

Terminal 1: `go run .`   Terminal 2: `Invoke-RestMethod "http://localhost:8080/debug/slow?seconds=8"`
Press Ctrl+C in terminal 1 within those 8 seconds. The log shows: not ready -> finishing in-flight
requests -> repository closed -> bye, and the slow request still completes.
`/ready` returns 503 as soon as shutdown starts. A second Ctrl+C kills the process at once.
