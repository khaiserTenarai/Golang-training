# Task 06 — External Package

```bash
go get github.com/google/uuid@latest   # adds to go.mod, creates go.sum
go run .
```
- `go.mod` records the dependency and version.
- `go.sum` stores checksums so builds are reproducible and tamper-proof.
- `go mod tidy` adds missing / removes unused dependencies.

Example output:
```
6f1c2e9a-...  Asha Rao     Engineering
...
Invalid ID "not-a-uuid" rejected: invalid UUID length: 10
```
