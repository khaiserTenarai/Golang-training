# Task 12 — Project Structure

```
12-project-structure/
├── go.mod                      module github.com/example/employeeapp
├── cmd/
│   └── employeeapp/main.go     entry point → go build ./cmd/employeeapp
├── internal/
│   ├── config/                 private: configuration loading
│   └── store/                  private: data storage
└── pkg/
    └── validator/              public: reusable validation helpers
```

| Directory | Purpose |
|-----------|---------|
| **cmd/** | One sub-folder per executable (`cmd/employeeapp`, later maybe `cmd/migrate`, `cmd/worker`). Each has a small `package main` that just wires things up. |
| **internal/** | Private application code. **The Go compiler enforces** that packages under `internal/` can only be imported by code rooted at the parent of `internal/`. Other modules get: `use of internal package ... not allowed`. Put business logic here so you can refactor freely. |
| **pkg/** | Library code that is **safe for external projects** to import. Treat it as a public API — changes here affect semantic versioning. (A convention, not compiler-enforced.) |

Run: `go run ./cmd/employeeapp`
