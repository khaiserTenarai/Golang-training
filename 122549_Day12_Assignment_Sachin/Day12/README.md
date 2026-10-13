# Day 12 - Production REST API Architecture

Same format as Day 11. Files 02-14 are single files (`go run <file>.go`)
and ask for their settings at startup. Folders `01_layered_architecture`
and `15_day12_mini_project` are multi-package projects (`go run .`
inside the folder; the mini project also needs `go mod tidy` and a
PostgreSQL database).

| File | Topic |
|------|-------|
| 01_layered_architecture | handler -> service -> repository as real packages |
| 02 | handler vs service |
| 03 | service vs repository (terminal app) |
| 04 | dependency injection |
| 05 | request-ID middleware |
| 06 | logging middleware |
| 07 | authentication (bearer token) |
| 08 | authorization (admin vs user) |
| 09 | validation layer |
| 10 | pagination |
| 11 | filtering |
| 12 | sorting |
| 13 | /health |
| 14 | /ready (checks a dependency you name) |
| 15_day12_mini_project | all of the above on PostgreSQL |

Note: written and hand-checked without a Go toolchain or database
available, so run `go vet` and try each file once before submitting.
