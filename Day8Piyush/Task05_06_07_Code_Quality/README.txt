============================================================
  Tasks 5, 6, 7 — Code Quality Tools
============================================================

TASK 5: Run gofmt
-----------------
  gofmt -d unformatted.go        # Shows formatting diff
  gofmt -w unformatted.go        # Formats the file in place
  gofmt -l .                     # Lists unformatted files

TASK 6: Run go vet
------------------
  go vet vet_issues.go           # Detects suspicious constructs
  go vet ./...                   # Vet all packages

  Common issues go vet catches:
  - Printf format mismatches
  - Unreachable code
  - Copying sync.Mutex
  - Wrong arg count in Printf

TASK 7: Static Analysis
-----------------------
  Install staticcheck:
    go install honnef.co/go/tools/cmd/staticcheck@latest

  Run:
    staticcheck static_analysis.go
    staticcheck ./...

  Common issues static analysis catches:
  - String concatenation in loops (use strings.Builder)
  - Compiling regex inside loops
  - Unnecessary type conversions
  - Empty branches
  - Deprecated function usage
