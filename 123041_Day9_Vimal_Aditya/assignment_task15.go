package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Employee struct {
	ID         int
	Name       string
	BaseSalary float64
	Age        int
}

type ProcessedResult struct {
	Employee   Employee
	Bonus      float64
	Tax        float64
	NetSalary  float64
	ProcessedBy int
	Status     string
}

// producer
func producer(jobs chan<- Employee, count int, wg *sync.WaitGroup) {
	defer wg.Done()

	names := []string{"Vimal", "Aditya", "Adi", "Alex", "Mason", "Connor", "Haytham", "Edward"}

	for i := 1; i <= count; i++ {
		emp := Employee{
			ID:         100 + i,
			Name:       names[(i-1)%len(names)],
			BaseSalary: 40000 + float64(i*5000),
			Age:        22 + (i % 15),
		}

		fmt.Printf("[Producer] Queuing Employee #%d: %s\n", emp.ID, emp.Name)
		jobs <- emp
		time.Sleep(100 * time.Millisecond)
	}

	close(jobs)
	fmt.Println("[Producer] Finished queuing all employees.")
}

// worker pool:
func worker(id int, jobs <-chan Employee, results chan<- ProcessedResult, cancelChan <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-cancelChan:
			fmt.Printf("   [Worker %d] Cancellation signal received. Shutting down worker.\n", id)
			return

		case emp, ok := <-jobs:
			if !ok {
				return
			}

			fmt.Printf("   [Worker %d] Processing Employee #%d (%s)...\n", id, emp.ID, emp.Name)

			time.Sleep(time.Duration(200+rand.Intn(300)) * time.Millisecond)

			bonus := emp.BaseSalary * 0.12
			tax := emp.BaseSalary * 0.18
			netSalary := emp.BaseSalary + bonus - tax

			result := ProcessedResult{
				Employee:    emp,
				Bonus:       bonus,
				Tax:         tax,
				NetSalary:   netSalary,
				ProcessedBy: id,
				Status:      "SUCCESS",
			}

			select {
			case results <- result:
				fmt.Printf("   [Worker %d] Completed Employee #%d\n", id, emp.ID)
			case <-time.After(2 * time.Second):
				fmt.Printf("   [Worker %d] Timeout sending results for Employee #%d\n", id, emp.ID)
			}
		}
	}
}

// consumer
func consumer(results <-chan ProcessedResult, done chan<- bool) {
	var totalPayroll float64
	processedCount := 0

	fmt.Println("\n=========================================================================")
	fmt.Printf("%-6s %-12s %-12s %-10s %-10s %-12s %-10s\n", "ID", "NAME", "BASE SALARY", "BONUS", "TAX", "NET SALARY", "WORKER")
	fmt.Println("=========================================================================")

	for res := range results {
		processedCount++
		totalPayroll += res.NetSalary

		fmt.Printf("%-6d %-12s %-12.2f %-10.2f %-10.2f %-12.2f Worker-%d\n",
			res.Employee.ID, res.Employee.Name, res.Employee.BaseSalary, res.Bonus, res.Tax, res.NetSalary, res.ProcessedBy)
	}

	fmt.Println("=========================================================================")
	fmt.Printf("Total Employees Processed: %d\n", processedCount)
	fmt.Printf("Total Net Payroll Output: $%.2f\n", totalPayroll)
	fmt.Println("=========================================================================")

	done <- true
}

// main controller
func main() {
	const totalEmployees = 8
	const numWorkers = 3

	jobs := make(chan Employee, 4)
	results := make(chan ProcessedResult, 8)
	cancelChan := make(chan struct{})
	consumerDone := make(chan bool)

	var producerWG sync.WaitGroup
	var workerWG sync.WaitGroup

	fmt.Println("=== STARTING CONCURRENT EMPLOYEE PROCESSING SYSTEM ===")

	producerWG.Add(1)
	go producer(jobs, totalEmployees, &producerWG)

	for w := 1; w <= numWorkers; w++ {
		workerWG.Add(1)
		go worker(w, jobs, results, cancelChan, &workerWG)
	}

	go consumer(results, consumerDone)

	go func() {
		workerWG.Wait()
		close(results)
	}()

	<-consumerDone

	fmt.Println("\n[System] All concurrent operations completed safely!")
}