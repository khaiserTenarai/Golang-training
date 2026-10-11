# Day 10 — Advanced Concurrency (Piyush)

No PostgreSQL needed. All tasks use only the Go standard library.

## How to Run

| Folder | Tasks | Command |
|--------|-------|---------|
| Task01_02_03_04_05_06_Sync_Primitives | 1–6 | `go run main.go` |
| Task07_08_09_10_Patterns | 7–10 | `go run main.go` |
| Task11_12_Context | 11, 12 | `go run main.go` |
| Task13_14_Deadlock_Race | 13, 14 | `go run -race main.go` |
| Task15_Mini_Project | 15 | `go run -race main.go` |

## Task Mapping

1. **sync.WaitGroup** → Task01–06 (demoWaitGroup)
2. **sync.Mutex** → Task01–06 (demoMutex)
3. **sync.RWMutex** → Task01–06 (demoRWMutex)
4. **sync.Once** → Task01–06 (demoOnce)
5. **sync.Map** → Task01–06 (demoSyncMap)
6. **Atomic counter** → Task01–06 (demoAtomic)
7. **Worker pool** → Task07–10 (demoWorkerPool)
8. **Fan-out** → Task07–10 (demoFanOut)
9. **Fan-in** → Task07–10 (demoFanIn)
10. **Concurrent pipeline** → Task07–10 (demoPipeline)
11. **Context cancellation** → Task11–12 (demoCancellation)
12. **Timeout handling** → Task11–12 (demoTimeout)
13. **Deadlock + fix** → Task13–14 (demoDeadlock)
14. **Race condition + fix** → Task13–14 (demoRace) — use `go run -race main.go`
15. **Mini Project** → Task15 (full system: worker pool, fan-out/in, timeout, race-safe, atomic counters)
