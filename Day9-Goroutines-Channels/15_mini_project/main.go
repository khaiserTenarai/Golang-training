package main

import (
	"fmt"
	"sync"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	BasePay   float64
	Hours     float64
	Processed bool
	NetSalary float64
}

// Producer: Sends employees to jobs channel
func employeeProducer(jobs chan<- Employee, employees []Employee) {
	for _, emp := range employees {
		jobs <- emp
	}
	close(jobs)
}

// Worker/Consumer: Processes employee calculations
func worker(id int, jobs <-chan Employee, results chan<- Employee, wg *sync.WaitGroup) {
	defer wg.Done()
	for emp := range jobs {
		// Calculate Overtime + Salary
		overtime := 0.0
		if emp.Hours > 40 {
			overtime = (emp.Hours - 40) * 1.5 * (emp.BasePay / 40)
		}
		emp.NetSalary = emp.BasePay + overtime
		emp.Processed = true

		fmt.Printf("[Worker %d] Processed ID: %d (%s) - Net: $%.2f\n", id, emp.ID, emp.Name, emp.NetSalary)
		results <- emp
	}
}

func main() {
	employeeList := []Employee{
		{ID: 1, Name: "Alice", BasePay: 3000, Hours: 40},
		{ID: 2, Name: "Bob", BasePay: 4000, Hours: 45},
		{ID: 3, Name: "Charlie", BasePay: 3500, Hours: 38},
		{ID: 4, Name: "Diana", BasePay: 5000, Hours: 50},
	}

	numJobs := len(employeeList)
	numWorkers := 2

	jobs := make(chan Employee, numJobs)
	results := make(chan Employee, numJobs)
	var workerWg sync.WaitGroup

	// Start Producer
	go employeeProducer(jobs, employeeList)

	// Start Worker Pool
	for w := 1; w <= numWorkers; w++ {
		workerWg.Add(1)
		go worker(w, jobs, results, &workerWg)
	}

	// Close results channel when workers finish
	go func() {
		workerWg.Wait()
		close(results)
	}()

	// Select monitoring channel completion
	done := make(chan bool)
	processedEmployees := []Employee{}

	go func() {
		for res := range results {
			processedEmployees = append(processedEmployees, res)
		}
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("\n--- Final Payroll Summary ---")
		for _, emp := range processedEmployees {
			fmt.Printf("ID: %d | Name: %-7s | Salary: $%.2f\n", emp.ID, emp.Name, emp.NetSalary)
		}
	case <-time.After(3 * time.Second):
		fmt.Println("Processing timed out!")
	}
}