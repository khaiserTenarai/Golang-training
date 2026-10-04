============================================================
  Task 15: Day 8 Mini Project — Tested Employee Service
============================================================

Features:
  - Employee CRUD (Add, Get, GetAll, Update, Delete)
  - Input validation (name, email, age, department, salary)
  - Salary calculations (net, annual, bonus, tax bracket)
  - Race-safe using sync.RWMutex
  - Comprehensive test suite

Commands:
---------

  Run the demo:
    go run main.go

  Run all tests with verbose output:
    go test -v ./service/...

  Run with race detector:
    go test -race -v ./service/...

  Measure test coverage:
    go test -cover ./service/...

  Generate coverage report (opens in browser):
    go test -coverprofile=coverage.out ./service/...
    go tool cover -html=coverage.out

  Run benchmarks:
    go test -bench=. -benchmem ./service/...

  Check formatting:
    gofmt -l .

  Run go vet:
    go vet ./...

Architecture:
  models/employee.go              → Data structure
  service/employee_service.go     → Business logic + validation
  service/employee_service_test.go → All tests + benchmarks
  main.go                         → Demo entry point
