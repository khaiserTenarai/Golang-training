# Day 9 — Goroutines & Channels (Piyush)

No PostgreSQL needed. All tasks use only the Go standard library.

## How to Run

| Folder | Tasks | Command |
|--------|-------|---------|
| Task01_02_03_Goroutine_Basics | 1, 2, 3 | `go run main.go` |
| Task04_05_06_07_08_Channels | 4, 5, 6, 7, 8 | `go run main.go` |
| Task09_10_Select | 9, 10 | `go run main.go` |
| Task11_12_Backpressure_ProducerConsumer | 11, 12 | `go run main.go` |
| Task13_Concurrent_File_Processor | 13 | `go run main.go` |
| Task14_Mini_Project | 14 | `go run main.go` |

## Task Mapping

1. **First goroutine** → Task01_02_03 (demoFirstGoroutine)
2. **Concurrent employee calculations** → Task01_02_03 (demoEmployeeCalculations)
3. **Goroutine lifecycle + anonymous goroutines** → Task01_02_03 (demoGoroutineLifecycle)
4. **Unbuffered channel** → Task04_05_06_07_08 (demoUnbuffered)
5. **Buffered channel** → Task04_05_06_07_08 (demoBuffered)
6. **Directional channels** → Task04_05_06_07_08 (demoDirectional)
7. **Channel closing** → Task04_05_06_07_08 (demoChannelClosing)
8. **Range over channel** → Task04_05_06_07_08 (demoRange)
9. **Select statement** → Task09_10 (demoSelect)
10. **Non-blocking operations** → Task09_10 (demoNonBlocking)
11. **Backpressure** → Task11_12 (demoBackpressure)
12. **Producer-Consumer** → Task11_12 (demoProducerConsumer)
13. **Concurrent file processor** → Task13 (worker pool pattern)
14. **Mini Project** → Task14 (full pipeline: producer → workers → consumer → report)
