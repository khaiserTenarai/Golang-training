# Day 9 - Goroutines & Channels

Same format as the earlier days: one `.go` file per question (the
question in a comment at the top, solution below), run individually
with `go run <filename>`. Every task reads its data from you at the
terminal instead of using hardcoded values.

```
01_first_goroutine.go
02_concurrent_employee_calculations.go
03_goroutine_lifecycle.go
04_unbuffered_channel.go
05_buffered_channel.go
06_directional_channels.go
07_channel_closing.go
08_range_over_channel.go
09_select.go
10_non_blocking_channel_ops.go
11_backpressure.go
12_producer_consumer.go
13_concurrent_file_processor.go
14_day9_mini_project.go
```

## Running them

```bash
go run 01_first_goroutine.go
```

Each one will prompt you for whatever it needs (a name, some numbers,
file paths, etc).

## A couple of things worth knowing as you read these

- **Why some files close a channel in their own little goroutine**
  (tasks 13 and 14): once several worker goroutines are all reading
  from the same jobs channel, nobody can safely call `close()` on the
  *results* channel until every worker is done sending to it. The
  trick used here is a `sync.WaitGroup` plus a tiny goroutine that
  just waits (`wg.Wait()`) and then closes the channel - that's what
  lets the main `for range` / `select` loop know when to stop instead
  of hanging forever.
- **Task 4 waits on a `done` channel at the end.** Without it, main()
  could print "Send completed." and exit before the receiver
  goroutine gets a chance to run its own `fmt.Println` - when main()
  returns, every other goroutine is killed immediately, finished or
  not. This same idea (something to wait on) shows up in several of
  the other files too, just using `sync.WaitGroup` instead.
- **Task 11 (backpressure)** is easiest to see with a small buffer
  size (like 1 or 2) and a consumer delay of a couple of seconds -
  you'll see "Producing: 3" print and then the program visibly pause
  before "Producing: 4", because the channel's buffer is full and the
  producer is stuck waiting for the slow consumer to make room.
- **Task 14** ties everything together: a producer goroutine, a pool
  of worker goroutines (you choose how many), channels connecting
  them, and a `select` with a timeout while main collects the results
  - a small, worker-pool-shaped version of patterns used in real
    concurrent programs.

## Note

Written and hand-checked (imports, brace/paren balance, and manually
tracing which goroutine starts first in the trickier files to rule out
deadlocks) without a local Go toolchain in this environment - please
run `go build ./...` yourself to confirm, though each file follows
standard, well-worn concurrency patterns.
