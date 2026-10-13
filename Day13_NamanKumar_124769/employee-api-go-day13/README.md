
# Day 13 — Configurable Employee API

## Run
```bash
go run ./cmd/api
```

The server listens on `:8080` by default. Development uses a development-only fallback token secret so the project runs locally without extra setup.

## Environment variables

| Variable | Purpose | Default |
|---|---|---|
| `APP_ENV` | `development`, `test`, or `production` | `development` |
| `PORT` | HTTP port | `8080` in development |
| `JWT_SECRET` | Token signing secret, kept separate from non-secret config | Development-only fallback |
| `LOG_LEVEL` | `debug`, `info`, `warn`, `error` | Environment-specific |
| `READ_TIMEOUT` | HTTP read timeout | `10s` |
| `SHUTDOWN_TIMEOUT` | Graceful shutdown timeout | `10s` |
| `TOKEN_TTL` | Access-token lifetime | `8h` |
| `FEATURE_EMPLOYEE_DELETE` | Production delete feature flag | `false` |

Production requires `JWT_SECRET` with at least 32 characters. Configure secrets through a secret manager or deployment environment, never commit them to source control.

## Environments

- Development: debug logging, port 8080, development-only secret fallback, delete enabled.
- Test: port 8081, shorter timeouts and token lifetime, delete disabled.
- Production: info logging, no secret fallback, delete disabled unless `FEATURE_EMPLOYEE_DELETE=true`.

Set an environment before starting:
```bash
APP_ENV=test go run ./cmd/api
```

Production example:
```bash
APP_ENV=production PORT=8080 JWT_SECRET='replace-with-a-long-random-secret-value' FEATURE_EMPLOYEE_DELETE=false go run ./cmd/api
```

## Demo token
```bash
curl -X POST http://localhost:8080/auth/token -H "Content-Type: application/json" -d '{"username":"admin","password":"admin123"}'
```

Demo credentials are for local practice only. Replace them with secure credential storage or an identity provider before production use.

## Endpoints
- `GET /health`
- `GET /ready`
- `POST /auth/token`
- `POST /api/employees`
- `GET /api/employees`
- `GET /api/employees/{id}`
- `PUT /api/employees/{id}`
- `DELETE /api/employees/{id}` when the delete feature flag is enabled

All `/api/*` routes require a bearer token. Mutations require the admin role.

List query parameters: `page`, `limit`, `department`, `search`, `sort_by`, and `order`.

## Task-to-code map
1. Environment variables: `internal/config/config.go`
2. Configuration struct: `config.Config`
3. Configuration validation: `config.Validate`
4. Defaults: `config.Load`
5. Development configuration: `APP_ENV=development`
6. Test configuration: `APP_ENV=test`
7. Production configuration: `APP_ENV=production`
8. Secrets separated from config: `config.Secrets`
9. Graceful shutdown: `internal/app/lifecycle.go`
10. SIGINT/SIGTERM: `signal.NotifyContext`
11. Context through service and repository methods: `internal/service/employee_service.go`, `internal/repository/employee_repository.go`
12. Startup initialization: `cmd/api/main.go`
13. Shutdown sequence: `internal/app/lifecycle.go`
14. Feature flag: `FEATURE_EMPLOYEE_DELETE` and `Config.EnableEmployeeDelete`
15. Mini project: the complete application composition in `cmd/api/main.go`

The repository remains in-memory so the project can run without external services.
