# 7. Perform static analysis

"Static analysis" just means checking code without actually running it -
`go vet` (question 6) is one static analysis tool built into Go, but it's
deliberately conservative (it only flags things it's very sure are bugs).
For deeper checks, the Go community mostly uses a separate tool called
`staticcheck`.

## Installing it

```
go install honnef.co/go/tools/cmd/staticcheck@latest
```

## Running it

```
staticcheck ./...
```

## What it catches here

`static_problem.go` computes `bonus` two different ways, but the first
result is immediately thrown away because `bonus` gets reassigned before
it's ever read. This is a real, working program - it compiles, `go vet`
says nothing about it - but it's almost certainly not what the author
meant to write. `staticcheck` flags this as an "ineffectual assignment":

```
static_problem.go:6:2: this value of bonus is never used (SA4006)
```

`static_fixed.go` just removes the dead line, keeping only the
calculation that's actually used.

I don't have `staticcheck` (or a Go toolchain at all) available in this
sandbox to run it for real, but this is a genuine, well-known staticcheck
finding - it's worth running `staticcheck ./...` alongside `go vet ./...`
on any real project, since they catch different classes of problems.
