// Task 1: Debug a Faulty Go Program — FIXED VERSION
// All 6 bugs from the buggy version have been identified and fixed.
// Run: go run main.go

package main

import (
	"fmt"
	"sync"
)

// Fix 1: Exported struct fields (capitalized) for external access
type Employee struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	employees := []Employee{
		{"Alice", 30, 50000},
		{"Bob", 25, 40000},
		{"Charlie", 35, 60000},
	}

	// Fix 2: Changed i <= len to i < len (off-by-one fix)
	for i := 0; i < len(employees); i++ {
		fmt.Println(employees[i].Name)
	}

	// Fix 3: Initialize the map using make()
	salaryMap := make(map[string]float64)
	salaryMap["Alice"] = 50000

	// Fix 4: Use == for comparison, not =
	x := 10
	if x == 10 {
		fmt.Println("x is 10")
	}

	// Fix 5: Pass loop variable as parameter to goroutine
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			fmt.Println("Value:", val) // Now prints 0, 1, 2 correctly
		}(i)
	}
	wg.Wait()

	// Fix 6: Check for zero before dividing
	total := 100
	count := 0
	if count != 0 {
		average := total / count
		fmt.Println("Average:", average)
	} else {
		fmt.Println("Cannot divide by zero. Count is 0.")
	}

	fmt.Println("Salary Map:", salaryMap)
	fmt.Println("\nAll bugs fixed successfully!")
}
