# Task 04 — Go Toolchain Practice

Run these inside this folder:

| Command | Purpose | Example output |
|---------|---------|----------------|
| `go run .` | Compiles to a temp dir and runs immediately. Great for development; leaves no binary behind. | `Sum of 1..10 = 55` |
| `go build -o app .` | Compiles the package into a standalone executable (`./app`). Used for releases/deployment. Cross-compile with `GOOS=linux GOARCH=arm64 go build`. | *(creates `app`)* |
| `go fmt ./...` | Rewrites source files into the official Go style (tabs, spacing, import ordering). Prints the names of files it changed. | *(no output = already formatted)* |
| `go vet ./...` | Static analysis that finds suspicious code that still compiles: wrong `Printf` verbs, unreachable code, copied locks, bad struct tags. | *(no output = no issues)* |
| `go test ./...` | Finds `*_test.go` files, runs `TestXxx` functions. `-v` for verbose, `-cover` for coverage. | `ok  gotoolchain  0.002s` |

### Try go vet catching a bug
Change `main.go` to `fmt.Printf("Sum = %d\n", "oops")` and run `go vet .`:
```
./main.go:6:2: fmt.Printf format %d has arg "oops" of wrong type string
```
