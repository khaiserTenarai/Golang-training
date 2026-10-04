# Golang Training — Naman Kumar

This repository contains my Go training assignments, organized by day.

## Project Structure

```text
Golang-training/
├── Day1NamanKumar/
│   └── go-day1/              # Go fundamentals and tooling (15 exercises)
├── Day2NamanKumar_124769/
│   └── day2/                 # Language basics: variables, types, loops, maps
├── Day3NamanKumar_124769/
│   └── employee-project/     # Employee modeling and reporting
├── Day4NamanKumar_124769/
│   └── employee-management/  # Layered employee management (model/service/utility)
├── Day5NamanKumar_124769/    # Repository-service architecture
├── Day6NamanKumar_124769/    # (add a short description)
├── Day7NamanKumar_124769/    # (add a short description)
├── Day8NamanKumar_124769/    # (add a short description)
└── Day9NamanKumar_124769/
    ├── task01_first_goroutine/
    ├── task02_concurrent_employee_calculations/
    ├── task03_goroutine_lifecycle/
    ├── task04_anonymous_goroutines/
    ├── task05_unbuffered_channel/
    ├── task06_buffered_channel/
    ├── task07_directional_channels/
    ├── task08_channel_closing/
    ├── task09_range_over_channel/
    ├── task10_select_statement/
    ├── task11_non_blocking_channel_ops/
    ├── task12_backpressure/
    ├── task13_producer_consumer/
    ├── task14_concurrent_file_processor/
    └── task15_miniproject_employee_processing/
```

## Daily Breakdown

| Day | Topic |
|-----|-------|
| Day 1 | Go fundamentals & tooling |
| Day 2 | Language basics (variables, types, loops, maps) |
| Day 3 | Employee modeling & reporting |
| Day 4 | Layered employee management (model/service/utility) |
| Day 5 | Repository-service architecture |
| Day 6 | _add topic_ |
| Day 7 | _add topic_ |
| Day 8 | _add topic_ |
| Day 9 | Goroutines, channels, select, producer-consumer, mini project |

## Running the Code

Each exercise folder has its own `go.mod`. To run an example:

```bash
cd Day1NamanKumar/go-day1/<folder-name>
go run main.go
```

For Day 9 tasks:

```bash
cd Day9NamanKumar_124769/<task-folder>
go run main.go
```

To run tests in a folder that has them:

```bash
go test ./...
```

## Notes

This repo is part of a Go training program. Each day's folder is self-contained with its own module.
