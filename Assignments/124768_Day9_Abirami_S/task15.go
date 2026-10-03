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

func producers(employees []Employee, ch chan<- Employee) {
	for _, employee := range employees {
		fmt.Println("Producing employee:", employee.Name)
		ch <- employee
	}
	close(ch)
}

func worker(id int, jobs <-chan Employee, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for employee := range jobs {
		fmt.Printf("Worker %d processing %s\n", id, employee.Name)
		time.Sleep(1 * time.Second)
		newSalary := employee.Salary + (employee.Salary * 10 / 100)
		result := fmt.Sprintf(
			"Worker %d: %s - Updated Salary: %.2f",
			id,
			employee.Name,
			newSalary,
		)
		results <- result
	}
}

func consumers(results <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for result := range results {
		fmt.Println("Result:", result)
	}
}

func main() {
	employees := []Employee{
		{ID: 1, Name: "Tom", Salary: 30000},
		{ID: 2, Name: "Priya", Salary: 35000},
		{ID: 3, Name: "Rahul", Salary: 40000},
		{ID: 4, Name: "Arun", Salary: 45000},
		{ID: 5, Name: "Divya", Salary: 50000},
	}
	jobs := make(chan Employee, 2)
	results := make(chan string, 5)

	var workerWg sync.WaitGroup
	var consumerWg sync.WaitGroup

	workerCount := 3

	workerWg.Add(workerCount)

	for i := 1; i <= workerCount; i++ {
		go worker(i, jobs, results, &workerWg)
	}

	consumerWg.Add(1)

	go consumers(results, &consumerWg)
	go producers(employees, jobs)
	workerWg.Wait()
	close(results)
	consumerWg.Wait()
	fmt.Println("All employee processing completed")
}
