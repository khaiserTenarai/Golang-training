# Day 11 - REST API with Go

One file per question (question in the header comment). Run a single
file with `go run <file>.go`. Each program asks you for what it needs
when it starts (port, starting employees, database details) and then
you talk to it with curl or a browser - the curl commands are in each
file's header comment.

Tasks 14 and 15 are folders because they use PostgreSQL: run
`go mod tidy` once inside the folder (downloads the driver), make sure
a database exists (e.g. `createdb employee_db`), then `go run .`.

| File | Topic |
|------|-------|
| 01 | basic HTTP server |
| 02 | GET /employee |
| 03 | GET /employee/{id} |
| 04 / 05 / 06 | POST / PUT / DELETE |
| 07 | correct status codes |
| 08 | JSON responses |
| 09 | JSON request bodies |
| 10 | path parameters |
| 11 | query parameters |
| 12 | logging middleware |
| 13 | validating employee data |
| 14_postgres_connection | connecting to PostgreSQL |
| 15_day11_mini_project | PostgreSQL-backed REST API, full CRUD |

Note: written and hand-checked (imports, braces, duplicate names)
without a Go toolchain or database available, so run `go vet` and try
each file once before submitting.
