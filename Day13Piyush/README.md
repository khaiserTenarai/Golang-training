# Day 13 — Configuration & Application Lifecycle (Piyush)

No PostgreSQL. In-memory store with 4 seed employees.

## Run
```bash
go mod tidy
go run main.go
```

## Override Config via Environment
```bash
APP_ENV=production FEATURE_AUTH=true LOG_LEVEL=info go run main.go
```

## Graceful Shutdown
Press `Ctrl+C` and watch the shutdown sequence.

## Task Mapping

| Task | What | File |
|------|------|------|
| 1 | Read env vars | config/config.go (getEnv) |
| 2 | Config struct | config/config.go (Config) |
| 3 | Validate config | config/config.go (Validate) |
| 4 | Config defaults | config/config.go (default values) |
| 5 | Dev config | config/environments.go (applyDevDefaults) |
| 6 | Test config | config/environments.go (applyTestDefaults) |
| 7 | Prod config | config/environments.go (applyProdDefaults) |
| 8 | Separate secrets | config/secrets.go |
| 9 | Graceful shutdown | main.go (server.Shutdown) |
| 10 | SIGINT/SIGTERM | main.go (signal.Notify) |
| 11 | Context through services | service/ + repository/ (all methods take ctx) |
| 12 | Startup sequence | main.go (Steps 1-8) |
| 13 | Shutdown sequence | main.go (Steps 1-4) |
| 14 | Feature flags | features/flags.go + handler checks |
| 15 | Mini Project | Everything combined |

## Architecture
```
main.go → startup/shutdown orchestration
├── config/         → Tasks 1-8 (config, envs, secrets)
├── features/       → Task 14 (feature flags)
├── middleware/      → Request-ID, logging, recovery, auth
├── handler/        → HTTP handlers (feature-flag aware)
├── service/        → Business logic (context-aware)
├── repository/     → Data access (context-aware)
└── models/         → Shared types
```
