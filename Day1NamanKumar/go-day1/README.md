# Day 1 — Go Fundamentals & Environment Setup

Each task is in its own folder (each is an independent Go module — `cd` into it and run).

| # | Folder | Task | Try it |
|---|--------|------|--------|
| 01 | `01-go-installation/` | Install Go, `go version`, `go env` (GOROOT/GOPATH) | read README |
| 02 | `02-hello-cloud-native/` | App name, version, Go version, environment | `APP_ENV=staging go run .` |
| 03 | `03-cli-calculator/` | + − × ÷ with input validation | `go run . 10 / 4` |
| 04 | `04-go-toolchain/` | run / build / fmt / vet / test explained | see README |
| 05 | `05-go-module/` | Module `employeeapp` + separate package | `go run .` |
| 06 | `06-external-package/` | `github.com/google/uuid` for employee IDs | `go run .` |
| 07 | `07-package-design/` | `main`, `employee`, `utils` packages | `go run .` |
| 08 | `08-go-documentation/` | GoDoc comments + generated docs | `go doc -all ./employee` |
| 09 | `09-versioning/` | Tags v1.0.0, v1.1.0, v1.1.1 + SemVer | `./create-tags.sh` |
| 10 | `10-environment-variables/` | APP_NAME / APP_PORT with defaults | `APP_PORT=9090 go run .` |
| 11 | `11-cli-employee-search/` | In-memory lookup by ID from CLI arg | `go run . 102` |
| 12 | `12-project-structure/` | `cmd/`, `internal/`, `pkg/` explained | `go run ./cmd/employeeapp` |
| 13 | `13-build-automation/` | Makefile + build.sh | `make all` |
| 14 | `14-troubleshooting/` | 2 intentional errors, diagnosis, fixes | `cd broken && go build ./...` |
| 15 | `15-mini-project-employee-management/` | Employee Management CLI | `make run` |

Tested with Go 1.22.
