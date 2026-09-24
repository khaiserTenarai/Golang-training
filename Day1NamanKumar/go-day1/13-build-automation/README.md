# Task 13 — Build Automation

## Makefile
```bash
make help    # list targets
make run     # go run ./cmd/app
make build   # bin/app with version embedded
make test    # go test -v -cover ./...
make fmt     # go fmt ./...
make vet     # go vet ./...
make all     # fmt → vet → test → build (use this before committing / in CI)
```

## Shell script (no make installed)
```bash
./build.sh run | build | test | fmt | vet | all | clean
```
On Windows, use Git Bash/WSL, or install make via `choco install make`.
