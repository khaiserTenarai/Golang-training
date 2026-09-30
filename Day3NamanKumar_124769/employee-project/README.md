# Employee Management System

A simple Go command-line application for managing employee records and
generating basic reports. Built as a Day 3 mini project to practice a
complete Git workflow: branching, committing, merging, tagging, and
collaborating through a remote.

## Features

- Add and list employees
- Generate a basic headcount/salary report
- Clean, dependency-free Go standard library implementation

## Project Structure

```
.
├── main.go        # Application entrypoint
├── employee.go    # Employee struct and related operations
├── report.go      # Reporting logic
├── go.mod         # Go module definition
└── README.md
```

## Requirements

- Go 1.22 or later

## Getting Started

Clone the repository and run the application:

```bash
git clone <repository-url>
cd employee-project
go run .
```

## Development Workflow

This project follows a feature-branch Git workflow:

1. Create a feature branch off `main` (e.g. `feature/employee`, `feature/report`)
2. Commit changes with clear, conventional messages
3. Open a pull request / merge back into `main`
4. Tag stable releases using annotated tags (e.g. `v1.0.0`)

## License

This project is for educational purposes.
