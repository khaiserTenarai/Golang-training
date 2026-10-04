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
	EmployeeName string
	OldSalary    float64
	Bonus        float64
	NewSalary    float64
}

func producer(employees chan<- Employee) {
	employeeList := []Employee{
		{ID: 101, Name: "Indu", Salary: 30000},
		{ID: 102, Name: "Rahul", Salary: 40000},
		{ID: 103, Name: "Priya", Salary: 50000},
		{ID: 104, Name: "Arun", Salary: 60000},
		{ID: 105, Name: "Anita", Salary: 35000},
		{ID: 106, Name: "Kiran", Salary: 45000},
	}

	for _, employee := range employeeList {
		fmt.Println("Producer: Sending", employee.Name)
		employees <- employee
		time.Sleep(300 * time.Millisecond)
	}

	close(employees)

	fmt.Println("Producer: Finished")
}

func worker(id int, employees <-chan Employee, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for employee := range employees {
		fmt.Println("Worker", id, "processing", employee.Name)

		bonus := employee.Salary * 0.10
		newSalary := employee.Salary + bonus

		result := Result{
			EmployeeName: employee.Name,
			OldSalary:    employee.Salary,
			Bonus:        bonus,
			NewSalary:    newSalary,
		}

		results <- result

		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("Worker", id, "finished")
}

func consumer(results <-chan Result) {
	for {
		select {
		case result, ok := <-results:
			if !ok {
				fmt.Println("Consumer: No more results")
				return
			}

			fmt.Printf(
				"Consumer: %s | Old Salary: %.2f | Bonus: %.2f | New Salary: %.2f\n",
				result.EmployeeName,
				result.OldSalary,
				result.Bonus,
				result.NewSalary,
			)

		case <-time.After(2 * time.Second):
			fmt.Println("Consumer: Waiting for results...")
		}
	}
}

func main() {
	employees := make(chan Employee)
	results := make(chan Result)

	var wg sync.WaitGroup

	go producer(employees)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, employees, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	consumer(results)

	fmt.Println("Employee processing completed")
}