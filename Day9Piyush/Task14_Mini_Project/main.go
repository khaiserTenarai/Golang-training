// Task 14: Day 9 Mini Project — Concurrent Employee Processing System
// Uses: Goroutines, Channels, Select, Producer, Consumer, Workers
// Run: go run main.go

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ============================================================
// Models
// ============================================================

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

type ProcessedEmployee struct {
	Employee
	AnnualSalary float64
	Tax          float64
	NetSalary    float64
	Bonus        float64
	Grade        string
	ProcessedBy  int
	Duration     time.Duration
}

// ============================================================
// Producer: Generates employee records
// ============================================================

func employeeProducer(employees []Employee, out chan<- Employee, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, emp := range employees {
		fmt.Printf("  [Producer] Sending employee: %s (ID: %d)\n", emp.Name, emp.ID)
		out <- emp
		time.Sleep(time.Duration(50+rand.Intn(100)) * time.Millisecond)
	}
}

// ============================================================
// Worker: Processes employee salary calculations
// ============================================================

func salaryWorker(id int, in <-chan Employee, out chan<- ProcessedEmployee, wg *sync.WaitGroup) {
	defer wg.Done()
	for emp := range in {
		start := time.Now()

		annual := emp.Salary * 12
		taxRate := getTaxRate(annual)
		tax := annual * taxRate
		net := annual - tax
		bonus := calculateBonus(emp.Salary)
		grade := getGrade(emp.Salary)

		// Simulate processing time
		time.Sleep(time.Duration(100+rand.Intn(200)) * time.Millisecond)

		result := ProcessedEmployee{
			Employee:     emp,
			AnnualSalary: annual,
			Tax:          tax,
			NetSalary:    net,
			Bonus:        bonus,
			Grade:        grade,
			ProcessedBy:  id,
			Duration:     time.Since(start),
		}

		fmt.Printf("  [Worker %d] Processed: %s | Net: %.0f | Grade: %s\n",
			id, emp.Name, net, grade)
		out <- result
	}
}

func getTaxRate(annual float64) float64 {
	switch {
	case annual <= 250000:
		return 0.0
	case annual <= 500000:
		return 0.05
	case annual <= 1000000:
		return 0.20
	default:
		return 0.30
	}
}

func calculateBonus(salary float64) float64 {
	if salary >= 70000 {
		return salary * 0.20
	} else if salary >= 50000 {
		return salary * 0.15
	} else if salary >= 30000 {
		return salary * 0.10
	}
	return salary * 0.05
}

func getGrade(salary float64) string {
	switch {
	case salary >= 80000:
		return "A"
	case salary >= 60000:
		return "B"
	case salary >= 40000:
		return "C"
	default:
		return "D"
	}
}

// ============================================================
// Consumer: Collects and aggregates results using select
// ============================================================

func resultConsumer(results <-chan ProcessedEmployee, done chan<- []ProcessedEmployee) {
	var all []ProcessedEmployee
	timeout := time.After(10 * time.Second)

	for {
		select {
		case result, ok := <-results:
			if !ok {
				// Channel closed, all results received
				done <- all
				return
			}
			all = append(all, result)
		case <-timeout:
			fmt.Println("  [Consumer] Timeout! Sending partial results.")
			done <- all
			return
		}
	}
}

// ============================================================
// Report Generator
// ============================================================

func generateReport(processed []ProcessedEmployee) {
	fmt.Println("\n╔══════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                    EMPLOYEE PROCESSING REPORT                                  ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Printf("║ %-4s %-12s %-12s %-10s %-10s %-10s %-6s %-6s ║\n",
		"ID", "Name", "Department", "Annual", "Tax", "Net", "Grade", "Worker")
	fmt.Println("╠══════════════════════════════════════════════════════════════════════════════════╣")

	var totalNet, totalTax, totalBonus float64
	for _, p := range processed {
		fmt.Printf("║ %-4d %-12s %-12s %-10.0f %-10.0f %-10.0f %-6s %-6d ║\n",
			p.ID, p.Name, p.Department, p.AnnualSalary, p.Tax, p.NetSalary, p.Grade, p.ProcessedBy)
		totalNet += p.NetSalary
		totalTax += p.Tax
		totalBonus += p.Bonus
	}

	fmt.Println("╠══════════════════════════════════════════════════════════════════════════════════╣")
	fmt.Printf("║ TOTAL: Employees=%-3d Net Salary=%-12.0f Tax=%-12.0f Bonus=%-10.0f ║\n",
		len(processed), totalNet, totalTax, totalBonus)
	fmt.Println("╚══════════════════════════════════════════════════════════════════════════════════╝")
}

// ============================================================
// Main — Orchestrates the Pipeline
// ============================================================

func main() {
	fmt.Println("===== DAY 9 MINI PROJECT: Concurrent Employee Processing System =====\n")

	employees := []Employee{
		{1, "Piyush", "Engineering", 75000},
		{2, "Alice", "Design", 55000},
		{3, "Bob", "Marketing", 45000},
		{4, "Charlie", "Engineering", 85000},
		{5, "Diana", "HR", 50000},
		{6, "Eve", "Finance", 65000},
		{7, "Frank", "Engineering", 90000},
		{8, "Grace", "Design", 48000},
		{9, "Hank", "Marketing", 42000},
		{10, "Ivy", "HR", 38000},
	}

	numWorkers := 3

	// Channels
	empChan := make(chan Employee, 5)        // Producer -> Workers
	resultChan := make(chan ProcessedEmployee, 10) // Workers -> Consumer
	doneChan := make(chan []ProcessedEmployee)      // Consumer -> Main

	fmt.Printf("Pipeline: %d employees → %d workers → consumer → report\n\n", len(employees), numWorkers)

	// Start consumer (uses select with timeout)
	go resultConsumer(resultChan, doneChan)

	// Start workers
	var workerWg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		workerWg.Add(1)
		go salaryWorker(i, empChan, resultChan, &workerWg)
	}

	// Start producer
	var producerWg sync.WaitGroup
	producerWg.Add(1)
	go employeeProducer(employees, empChan, &producerWg)

	// Close empChan when producer is done
	go func() {
		producerWg.Wait()
		close(empChan)
		fmt.Println("\n  [System] Producer finished, employee channel closed.")
	}()

	// Close resultChan when all workers are done
	go func() {
		workerWg.Wait()
		close(resultChan)
		fmt.Println("  [System] All workers finished, results channel closed.")
	}()

	// Wait for consumer to finish and get results
	processed := <-doneChan

	// Generate final report
	generateReport(processed)

	fmt.Println("\nConcurrent Employee Processing System completed!")
}
