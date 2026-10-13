# Day 10 - Advanced Concurrency

One `.go` file per question (question in the header comment, solution
below). Questions that need two files get their own folder. Every
program asks for its input at the terminal.

Run any single file with `go run <file>.go`.

## Files

- 01_waitgroup.go, 02_mutex.go, 03_rwmutex.go
- 04_sync_once.go, 05_sync_map.go, 06_atomic_counter.go
- 07_worker_pool.go, 08_fan_out.go, 09_fan_in.go, 10_concurrent_pipeline.go
- 11_context_cancellation.go, 12_timeout_handling.go
- 13_deadlock/ - `deadlock_example.go` (crashes on purpose) and `deadlock_fixed.go`
- 14_race_condition/ - `race_buggy.go` and `race_fixed.go`
- 15_day10_mini_project.go

## Things worth trying

- `go run deadlock_example.go` crashes with
  `fatal error: all goroutines are asleep - deadlock!` - that is the
  expected result, it is the demo.
- `go run -race race_buggy.go` prints a `DATA RACE` report. Run
  `go run -race race_fixed.go` afterwards and it stays quiet. Use a
  fairly large goroutine count (e.g. 1000) for the buggy one.
- In the mini project, pick a timeout shorter than
  (jobs / workers) x 0.2s to see the timeout path trigger.

## Note

Written and hand-checked without a Go toolchain (no Go or internet in
the sandbox where this was prepared), so please run `go vet` and the
files once on your machine before submitting.
