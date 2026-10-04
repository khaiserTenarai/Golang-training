============================================================
  Tasks 9, 10, 11, 12, 13 — Testing
============================================================

Task  9: Unit tests for employee validation  → employee_test.go
Task 10: Salary calculation tests            → employee_test.go
Task 11: Measure test coverage               → commands below
Task 12: Table-driven tests                  → employee_test.go (TestXxx_TableDriven)
Task 13: Benchmarks                          → employee_benchmark_test.go

Commands to Run:
----------------

  Run all tests:
    go test -v

  Run specific test:
    go test -v -run TestValidateEmployee_TableDriven

  Task 11 — Measure coverage:
    go test -cover

  Coverage with report:
    go test -coverprofile=coverage.out
    go tool cover -html=coverage.out   (opens browser)
    go tool cover -func=coverage.out   (prints per-function %)

  Task 13 — Run benchmarks:
    go test -bench=. -benchmem

  Run specific benchmark:
    go test -bench=BenchmarkCalculateNetSalary -benchmem

  Race detection during tests:
    go test -race -v
