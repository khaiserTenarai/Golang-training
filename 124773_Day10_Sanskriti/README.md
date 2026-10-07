# Go Day 10 - Concurrency Programs

This folder contains 15 simple Go programs for concurrency practice.

1. WaitGroup
2. Mutex
3. RWMutex
4. sync.Once
5. sync.Map
6. Atomic Counter
7. Worker Pool
8. Fan-Out
9. Fan-In
10. Concurrent Pipeline
11. Context Cancellation
12. Timeout Handling
13. Deadlock and Fix
14. Race Condition and Fix
15. Day 10 Mini Project

## Running a program

Open the program folder and run:

    go run main.go

Example:

    cd 01_waitgroup
    go run main.go

## Race detector

For programs 14 and 15:

    go run -race main.go

The comments in each program explain the important parts.
