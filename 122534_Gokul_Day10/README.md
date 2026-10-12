One .go file per question (question in the header comment, solution below). Questions that need two files get their own folder. Every program asks for its input at the terminal.

Files

01_waitgroup.go
02_mutex.go
03_rwmutex.go
04_sync_once.go
05_sync_map.go
06_atomic_counter.go
07_worker_pool.go
08_fan_out.go
09_fan_in.go
10_concurrent_pipeline.go
11_context_cancellation.go
12_timeout_handling.go
13_deadlock/
deadlock_example.go (crashes on purpose)
deadlock_fixed.go
14_race_condition/
race_buggy.go
race_fixed.go
15_day10_mini_project.go

Things worth trying

Run go run deadlock_example.go. It crashes with fatal error: all goroutines are asleep - deadlock!. This is the expected result of the demo.
Run go run -race race_buggy.go. It prints a DATA RACE report. Then run go run -race race_fixed.go, which should complete without a race report. Use a fairly large goroutine count, such as 1000, for the buggy version.
In the mini project, pick a timeout shorter than (jobs / workers) × 0.2 seconds to trigger the timeout path.

Note

Written and hand-checked without a Go toolchain (no Go or internet in the sandbox where this was prepared), so please run go vet and test the files on your machine before submitting.