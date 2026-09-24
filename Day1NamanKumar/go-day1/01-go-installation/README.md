# Task 01 — Go Installation

## Install
- **Linux**: download from https://go.dev/dl, then
  `sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.22.x.linux-amd64.tar.gz`
  and add `export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin` to `~/.bashrc`.
- **macOS**: `brew install go` (or the .pkg installer).
- **Windows**: run the .msi installer from go.dev/dl.

## Verify
```bash
$ go version
go version go1.22.2 linux/amd64
```

## go env
```bash
$ go env GOROOT GOPATH
/usr/lib/go-1.22
/root/go
```

| Variable | Meaning |
|----------|---------|
| **GOROOT** | Where the Go SDK itself is installed (compiler, standard library, tools). Set automatically; rarely changed. |
| **GOPATH** | Your Go workspace. Holds downloaded module cache (`$GOPATH/pkg/mod`) and binaries from `go install` (`$GOPATH/bin`). Defaults to `$HOME/go`. With Go modules, your project code can live anywhere. |

Other useful vars: `GOOS`, `GOARCH` (target platform), `GOPROXY` (module download proxy), `GOMODCACHE`, `GOBIN`.
