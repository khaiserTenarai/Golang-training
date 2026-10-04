package main

import (
	"fmt"
)

var employees []string

func addEmployee(name string) {
	employees = append(employees, name)
}

func main() {
	var name string

	fmt.Scan(&name)

	addEmployee(name)

	fmt.Println(employees)

	for i := 0; i < len(employees); i++ {
		fmt.Println(employees[i])
	}
}

// go run code_review.go

// Run program

// gofmt -w .

// Format code

// go vet ./...

// Check suspicious code

// go test ./...

// Run all tests

// go test -v

// Run tests with details

// go test -cover

// Measure coverage

// go test -coverprofile=coverage.out

// Create coverage report

// go tool cover -html=coverage.out

// Run benchmarks

// go test -bench=.

// Detect race conditions

// go run -race race.go

// Static analysis

// staticcheck ./...