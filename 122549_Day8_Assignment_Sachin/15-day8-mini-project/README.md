# Day 8 Mini Project — Tested Employee Service

This folder is a small but genuinely tested Employee Service:
`service.go` holds the logic, `service_test.go` holds the tests, and
`main.go` is an interactive program that reads real input from the user
and uses the same service underneath.

## 1. Unit tests

```
go test -v ./...
```

Runs every `TestXxx` function in `service_test.go`:
`TestAddAndGet`, `TestGetNotFound`, `TestDelete`, `TestDeleteNotFound`,
`TestCount`, and `TestConcurrentAdd`. `-v` prints each test name and
PASS/FAIL individually instead of just a summary line.

## 2. Coverage

```
go test -cover ./...
```

Prints something like:

```
ok      employeeservice   0.010s  coverage: 92.3% of statements
```

For a full breakdown of exactly which lines aren't covered:

```
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 3. Formatting

```
gofmt -l .
```

Lists any file that isn't formatted according to Go's standard style. No
output means everything is already properly formatted. (`gofmt -w .`
would rewrite any file that needs it.)

## 4. Static analysis

```
go vet ./...
```

and, for deeper checks:

```
go install honnef.co/go/tools/cmd/staticcheck@latest
staticcheck ./...
```

## 5. Race detection

```
go test -race ./...
```

This is exactly what `TestConcurrentAdd` exists for - it fires 50
goroutines at `Add()` at the same time. Because `EmployeeService` protects
its internal map with a `sync.Mutex` (see `service.go`), this should pass
cleanly with no `WARNING: DATA RACE` output, and `Count()` should always
correctly report 50.

To see this actually matter, try temporarily removing the
`s.mu.Lock()` / `s.mu.Unlock()` lines from `Add()` and re-run
`go test -race ./...` - it should then report a data race, the same kind
shown in question 8's `race_buggy.go`.

## Running the program itself

```
go run .
```

Gives you an interactive menu to add, get, delete, and count employees -
all backed by the exact same `EmployeeService` that the tests above are
checking.

I don't have a Go toolchain available in this sandbox to actually run
these 5 commands and paste real output, but every command above is exactly
what you'd run, and the code is written specifically so each one has
something meaningful to check (a real mutex to race-test, a real untested
edge case to notice in coverage, etc).
