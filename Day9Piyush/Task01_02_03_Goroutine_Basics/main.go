// Tasks 1, 2, 3: Goroutine Basics
// Task 1 — Create your first goroutine
// Task 2 — Run multiple employee calculations concurrently
// Task 3 — Demonstrate goroutine life cycle + anonymous goroutines
// Run: go run main.go

package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 1: Create Your First Goroutine

func greet(name string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("[Goroutine] Hello, %s! (iteration %d)\n", name, i+1)
		time.Sleep(100 * time.Millisecond)
	}
}

func demoFirstGoroutine() {
	fmt.Println("========== TASK 1: First Goroutine ==========")

	go greet("Piyush") // Runs concurrently in a separate goroutine

	// Main goroutine continues
	for i := 0; i < 3; i++ {
		fmt.Printf("[Main] Working... (iteration %d)\n", i+1)
		time.Sleep(150 * time.Millisecond)
	}
	time.Sleep(200 * time.Millisecond)
	fmt.Println()
}

// Task 2: Run Multiple Employee Calculations Concurrently

type Employee struct {
	Name   string
	Salary float64
}

func calculateAnnualSalary(emp Employee, wg *sync.WaitGroup) {
	defer wg.Done()
	annual := emp.Salary * 12
	tax := annual * 0.20
	net := annual - tax
	fmt.Printf("  %s: Monthly=%.0f | Annual=%.0f | Tax=%.0f | Net=%.0f\n",
		emp.Name, emp.Salary, annual, tax, net)
}

func demoEmployeeCalculations() {
	fmt.Println("========== TASK 2: Concurrent Employee Calculations ==========")

	employees := []Employee{
		{"Piyush", 60000},
		{"Alice", 55000},
		{"Bob", 45000},
		{"Charlie", 70000},
		{"Diana", 50000},
	}

	var wg sync.WaitGroup
	for _, emp := range employees {
		wg.Add(1)
		go calculateAnnualSalary(emp, &wg)
	}
	wg.Wait()
	fmt.Println("All calculations completed!\n")
}

// Task 3: Goroutine Life Cycle + Anonymous Goroutines

func demoGoroutineLifecycle() {
	fmt.Println("========== TASK 3: Goroutine Life Cycle ==========")

	var wg sync.WaitGroup

	// Named function goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("  [Named] Goroutine started")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("  [Named] Goroutine doing work...")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("  [Named] Goroutine finished")
	}()

	// Task 3.1: Anonymous goroutine with parameter
	wg.Add(1)
	go func(msg string) {
		defer wg.Done()
		fmt.Printf("  [Anonymous] Received message: %s\n", msg)
		time.Sleep(150 * time.Millisecond)
		fmt.Println("  [Anonymous] Processing complete")
	}("Hello from main!")

	// Multiple anonymous goroutines in a loop
	fmt.Println("\n  Launching 5 anonymous worker goroutines:")
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("    Worker %d: started\n", id)
			time.Sleep(time.Duration(id*50) * time.Millisecond)
			fmt.Printf("    Worker %d: done\n", id)
		}(i) // Pass loop variable to avoid closure bug
	}

	wg.Wait()
	fmt.Println("All goroutines completed their lifecycle.\n")
}

func main() {
	demoFirstGoroutine()
	demoEmployeeCalculations()
	demoGoroutineLifecycle()
}
