package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a WaitGroup to wait for the goroutine to finish
	var wg sync.WaitGroup

	// Employee details
	name := "John"
	salary := 50000

	// Add one goroutine to the WaitGroup
	wg.Add(1)

	// Create an anonymous goroutine
	go func() {
		// Mark the goroutine as completed when it finishes
		defer wg.Done()

		// Calculate 10% bonus
		bonus := salary * 10 / 100

		// Calculate total salary
		totalSalary := salary + bonus

		// Display employee details
		fmt.Println("Employee:", name)
		fmt.Println("Basic Salary:", salary)
		fmt.Println("Bonus:", bonus)
		fmt.Println("Total Salary:", totalSalary)
	}()

	// Wait for the goroutine to complete
	wg.Wait()
}
