# 13. Create benchmarks

A benchmark measures how fast a piece of code runs, instead of just
whether it gives the right answer. In Go, a benchmark function starts with
`Benchmark` (instead of `Test`) and takes a `*testing.B` instead of a
`*testing.T`.

## Command

```
go test -bench=.
```

## What `b.N` means

Go doesn't know in advance how many times to run `factorial(10)` to get a
reliable timing, so it runs the loop with an increasing `b.N` (1, then
more, then more again) until the total time is long enough to measure
accurately, and then reports the average per call.

## Expected kind of output

```
BenchmarkFactorial-8   50000000     25.3 ns/op
```

- `50000000` - how many times the loop actually ran
- `25.3 ns/op` - average nanoseconds per single call to `factorial(10)`

I don't have a Go toolchain in this sandbox to produce a real number here,
but that's the shape of what you'd see - the exact numbers depend on the
machine it's run on, which is normal for benchmarks (they're for comparing
"before vs after" a code change on the same machine, not for treating the
raw number as an absolute fact).
