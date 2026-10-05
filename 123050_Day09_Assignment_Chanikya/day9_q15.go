package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

type Result struct {
	EmployeeName string
	Salary       float64
	WorkerID     int
}

// Producer sends employees to the employee channel.
func producer(employees []Employee, employeeChannel chan<- Employee) {
	for _, employee := range employees {
		fmt.Println("Producer: sending", employee.Name)

		employeeChannel <- employee
	}

	close(employeeChannel)

	fmt.Println("Producer: finished")
}

// Worker receives employees and processes them.
func worker(
	workerID int,
	employeeChannel <-chan Employee,
	resultChannel chan<- Result,
) {
	for employee := range employeeChannel {
		fmt.Printf(
			"Worker %d: processing %s\n",
			workerID,
			employee.Name,
		)

		finalSalary := employee.Salary + 5000

		time.Sleep(500 * time.Millisecond)

		resultChannel <- Result{
			EmployeeName: employee.Name,
			Salary:       finalSalary,
			WorkerID:     workerID,
		}
	}
}

// Consumer receives and displays the results.
func consumer(resultChannel <-chan Result, done chan<- bool) {
	for {
		select {
		case result, ok := <-resultChannel:
			if !ok {
				done <- true
				return
			}

			fmt.Printf(
				"Consumer: %s final salary = %.2f (Worker %d)\n",
				result.EmployeeName,
				result.Salary,
				result.WorkerID,
			)

		case <-time.After(3 * time.Second):
			fmt.Println("Consumer: timeout")
			done <- true
			return
		}
	}
}

func main() {
	employees := []Employee{
		{ID: 1, Name: "Amit", Salary: 50000},
		{ID: 2, Name: "Rahul", Salary: 60000},
		{ID: 3, Name: "Suresh", Salary: 70000},
		{ID: 4, Name: "Kiran", Salary: 55000},
		{ID: 5, Name: "Arun", Salary: 65000},
	}

	employeeChannel := make(chan Employee, 2)
	resultChannel := make(chan Result, 2)
	done := make(chan bool)

	// Start producer
	go producer(employees, employeeChannel)

	// Start 3 workers
	for workerID := 1; workerID <= 3; workerID++ {
		go worker(workerID, employeeChannel, resultChannel)
	}

	// Start consumer
	go consumer(resultChannel, done)

	// Wait for consumer
	<-done

	close(resultChannel)

	fmt.Println("Employee processing completed")
}
