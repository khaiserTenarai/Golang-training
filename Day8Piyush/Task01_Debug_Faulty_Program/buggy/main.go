// Task 1: Debug a Faulty Go Program — BUGGY VERSION
// This program has 6 intentional bugs. Find and fix them!
// Run: go run main.go

package main

import (
	"fmt"
)

// Bug 1: Struct field names start with lowercase — cannot be accessed outside package
type employee struct {
	name   string
	age    int
	salary float64
}

func main() {
	employees := []employee{
		{"Alice", 30, 50000},
		{"Bob", 25, 40000},
		{"Charlie", 35, 60000},
	}

	// Bug 2: Off-by-one error — i <= len causes index out of range
	for i := 0; i <= len(employees); i++ {
		fmt.Println(employees[i].name)
	}

	// Bug 3: Nil map — writing to a nil map causes panic
	var salaryMap map[string]float64
	salaryMap["Alice"] = 50000

	// Bug 4: Wrong comparison — using = instead of == (won't compile)
	x := 10
	if x = 10 {
		fmt.Println("x is 10")
	}

	// Bug 5: Goroutine closure captures loop variable by reference
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println("Value:", i) // Always prints 3
		}()
	}

	// Bug 6: Division by zero
	total := 100
	count := 0
	average := total / count
	fmt.Println("Average:", average)

	fmt.Println(salaryMap)
}
