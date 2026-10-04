# Day 8 — Testing, Debugging & Code Quality (Piyush)

No PostgreSQL needed. All tasks use only the Go standard library.

## How to Run Each Task

| Folder | Tasks | Command |
|--------|-------|---------|
| Task01_Debug_Faulty_Program/buggy | 1 | `go run main.go` (observe bugs) |
| Task01_Debug_Faulty_Program/fixed | 1 | `go run main.go` (bugs fixed) |
| Task02_03_04_Logging | 2, 3, 4 | `go run main.go` |
| Task05_06_07_Code_Quality | 5 | `gofmt -d unformatted.go` |
| Task05_06_07_Code_Quality | 6 | `go vet vet_issues.go` |
| Task05_06_07_Code_Quality | 7 | `staticcheck static_analysis.go` |
| Task08_Race_Detection | 8 | `go run -race main.go` |
| Task09_10_11_12_13_Testing | 9,10,12 | `go test -v` |
| Task09_10_11_12_13_Testing | 11 | `go test -cover` |
| Task09_10_11_12_13_Testing | 13 | `go test -bench=. -benchmem` |
| Task14_Code_Review | 14 | Read bad_code.go + review_report.txt |
| Task15_Mini_Project | 15 | `go run main.go` |
| Task15_Mini_Project | 15 | `go test -v -race -cover ./service/...` |

## Task Mapping

1. **Debug a faulty program** → Task01 (buggy/ and fixed/)
2. **Application logging** → Task02_03_04 (demoBasicLogging)
3. **Log levels** → Task02_03_04 (demoLogLevels)
4. **Structured logs** → Task02_03_04 (demoStructuredLogging)
5. **Run gofmt** → Task05_06_07 (unformatted.go)
6. **Run go vet** → Task05_06_07 (vet_issues.go)
7. **Static analysis** → Task05_06_07 (static_analysis.go)
8. **Race detection** → Task08 (main.go with -race flag)
9. **Unit tests** → Task09_10_11_12_13 (employee_test.go)
10. **Salary tests** → Task09_10_11_12_13 (employee_test.go)
11. **Coverage** → Task09_10_11_12_13 (`go test -cover`)
12. **Table-driven tests** → Task09_10_11_12_13 (TestXxx_TableDriven)
13. **Benchmarks** → Task09_10_11_12_13 (employee_benchmark_test.go)
14. **Code review** → Task14 (bad_code.go + review_report.txt)
15. **Mini Project** → Task15 (full Employee Service with all quality checks)
