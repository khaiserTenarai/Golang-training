package main

import (
	"fmt"
	"sync"
	"time"
)

// Employee represents an employee
type Employee struct {
	ID     int
	Name   string
	Salary float64
}

// Producer sends employees into the jobs channel
func producer(jobs chan<- Employee) {

	// Create employee data
	employees := []Employee{
		{1, "Abby", 30000},
		{2, "Rahul", 35000},
		{3, "Priya", 40000},
		{4, "John", 45000},
		{5, "Sara", 50000},
	}

	// Send each employee to the jobs channel
	for _, employee := range employees {

		fmt.Println("Producer: Sending", employee.Name)

		jobs <- employee
	}

	// Close the jobs channel after sending all employees
	close(jobs)

	fmt.Println("Producer: All employees sent")
}

// Worker processes employee jobs
func worker(id int, jobs <-chan Employee, results chan<- string, wg *sync.WaitGroup) {

	// Tell WaitGroup that this worker is completed
	defer wg.Done()

	// Receive employees from the jobs channel
	for employee := range jobs {

		fmt.Println("Worker", id, "processing", employee.Name)

		// Simulate employee processing
		time.Sleep(500 * time.Millisecond)

		// Calculate 10% bonus
		bonus := employee.Salary * 10 / 100

		// Calculate total salary
		totalSalary := employee.Salary + bonus

		// Create result message
		result := fmt.Sprintf(
			"%s - Total Salary: %.2f",
			employee.Name,
			totalSalary,
		)

		// Send result to the results channel
		results <- result
	}
}

// Consumer receives processed results
func consumer(results <-chan string, done <-chan bool) {

	// Keep receiving results
	for {

		select {

		// Receive processed employee result
		case result := <-results:

			fmt.Println("Consumer:", result)

		// Receive completion signal
		case <-done:

			fmt.Println("Consumer: Processing completed")

			return
		}
	}
}

func main() {

	// Create jobs channel
	jobs := make(chan Employee)

	// Create results channel
	results := make(chan string)

	// Create channel to signal completion
	done := make(chan bool)

	// Create WaitGroup for workers
	var wg sync.WaitGroup

	// Number of workers
	workerCount := 3

	// Add workers to WaitGroup
	wg.Add(workerCount)

	// Start producer goroutine
	go producer(jobs)

	// Start worker goroutines
	for i := 1; i <= workerCount; i++ {

		go worker(i, jobs, results, &wg)
	}

	// Close results after all workers finish
	go func() {

		// Wait for all workers to complete
		wg.Wait()

		// Tell consumer that processing is completed
		done <- true

		// Close results channel
		close(results)

	}()

	// Start consumer
	consumer(results, done)

	// Program completed
	fmt.Println("Employee processing system completed")
}
