# 6. Run go vet

`go vet` checks code that compiles fine, but is still likely wrong. A
classic example is a `Printf`-style format string that doesn't match its
arguments - the compiler doesn't catch this (it's just a string and some
values to it), but `go vet` specifically understands `Printf` verbs and
will flag it.

`vet_problem.go` uses `%d` (expects a number) for `name`, which is a
string.

## Command

```
go vet ./...
```

## Expected output

```
./vet_problem.go:11:2: Printf format %d has arg name of wrong type string
```

`vet_fixed.go` is the corrected version - `%d` changed to `%s` for the
name. Running `go vet` against it reports no issues.

I don't have a Go toolchain in this sandbox to actually run this, but this
is genuinely one of the most common things `go vet` catches in real
projects, so it's worth running before every commit.
