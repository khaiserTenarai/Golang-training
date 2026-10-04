// Task 6: Run go vet — This file has issues that go vet will catch.
// Run: go vet vet_issues.go
// Note: Some issues will prevent compilation; they demonstrate what go vet detects.

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Issue 1: Printf format mismatch — %d expects int but gets string
	name := "Piyush"
	fmt.Printf("Employee ID: %d\n", name)

	// Issue 2: Unreachable code after return
	result := calculate()
	fmt.Println("Result:", result)

	// Issue 3: Copying a sync.Mutex (go vet catches this)
	var mu sync.Mutex
	mu2 := mu // copying a lock
	_ = mu2

	// Issue 4: Wrong number of args in Printf
	age := 25
	fmt.Printf("Name: %s, Age: %d, Salary: %f\n", name, age)
}

func calculate() int {
	return 42
	fmt.Println("This line is unreachable") // go vet flags this
	return 0
}
