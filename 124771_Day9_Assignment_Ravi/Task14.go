package main

import (
	"fmt"
	"sync"
	"time"
)

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

type Result struct {
	EmployeeID int
	Name       string
	Salary     float64
	Bonus      float64
	Total      float64
}

// Producer creates employees and sends them to the jobs channel.
func producer(employees []Employee, jobs chan<- Employee) {
	defer close(jobs)

	for _, emp := range employees {
		fmt.Printf("Producer: sending %s\n", emp.Name)
		jobs <- emp
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Println("Producer: finished")
}

// Worker processes employees from the jobs channel.
func worker(id int, jobs <-chan Employee, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case emp, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: no more jobs\n", id)
				return
			}

			fmt.Printf("Worker %d: processing %s\n", id, emp.Name)

			// Simulate employee calculation
			time.Sleep(500 * time.Millisecond)

			bonus := emp.Salary * 0.10
			total := emp.Salary + bonus

			result := Result{
				EmployeeID: emp.ID,
				Name:       emp.Name,
				Salary:     emp.Salary,
				Bonus:      bonus,
				Total:      total,
			}

			results <- result

		case <-time.After(2 * time.Second):
			fmt.Printf("Worker %d: timeout, stopping\n", id)
			return
		}
	}
}

// Consumer receives processed employee results.
func consumer(results <-chan Result) {
	for {
		select {
		case result, ok := <-results:
			if !ok {
				fmt.Println("Consumer: no more results")
				return
			}

			fmt.Printf(
				"Consumer: %s | Salary: %.2f | Bonus: %.2f | Total: %.2f\n",
				result.Name,
				result.Salary,
				result.Bonus,
				result.Total,
			)

		case <-time.After(3 * time.Second):
			fmt.Println("Consumer: timeout")
			return
		}
	}
}

func main() {
	employees := []Employee{
		{1, "Alice", 50000},
		{2, "Bob", 60000},
		{3, "Charlie", 55000},
		{4, "David", 70000},
		{5, "Eva", 65000},
		{6, "Frank", 75000},
	}

	// Channels
	jobs := make(chan Employee, 2)
	results := make(chan Result, 2)

	// Start producer
	go producer(employees, jobs)

	// Start workers
	var wg sync.WaitGroup

	workerCount := 3

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Close results after all workers finish.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Consumer runs in the main goroutine.
	consumer(results)

	fmt.Println("All employee processing completed")
}
