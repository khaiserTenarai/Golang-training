package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

type ProcessedResult struct {
	EmployeeID int
	Name       string
	Bonus      float64
	Tax        float64
	NetPay     float64
}

func generateEmployees(count int) []Employee {
	employees := make([]Employee, 0, count)
	for i := 1; i <= count; i++ {
		employees = append(employees, Employee{
			ID:     i,
			Name:   fmt.Sprintf("Employee-%d", i),
			Salary: float64(30000 + rand.Intn(70000)),
		})
	}
	return employees
}

func producer(employees []Employee, jobs chan<- Employee) {
	for _, e := range employees {
		jobs <- e
	}
	close(jobs)
}

func worker(id int, jobs <-chan Employee, results chan<- ProcessedResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for e := range jobs {
		time.Sleep(10 * time.Millisecond)
		bonus := e.Salary * 0.10
		tax := e.Salary * 0.20
		net := e.Salary + bonus - tax
		results <- ProcessedResult{
			EmployeeID: e.ID,
			Name:       e.Name,
			Bonus:      bonus,
			Tax:        tax,
			NetPay:     net,
		}
	}
}

func consumer(results <-chan ProcessedResult, done chan<- []ProcessedResult) {
	var collected []ProcessedResult
	for r := range results {
		collected = append(collected, r)
	}
	done <- collected
}

func main() {
	employees := generateEmployees(20)

	jobs := make(chan Employee, len(employees))
	results := make(chan ProcessedResult, len(employees))
	done := make(chan []ProcessedResult)

	const workerCount = 4
	var wg sync.WaitGroup

	go producer(employees, jobs)

	for w := 1; w <= workerCount; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	go consumer(results, done)

	go func() {
		wg.Wait()
		close(results)
	}()

	timeout := time.After(5 * time.Second)

	select {
	case finalResults := <-done:
		fmt.Println("Processed employees:", len(finalResults))
		var totalNetPay float64
		for _, r := range finalResults {
			fmt.Printf("ID: %d | %s | Bonus: %.2f | Tax: %.2f | Net Pay: %.2f\n",
				r.EmployeeID, r.Name, r.Bonus, r.Tax, r.NetPay)
			totalNetPay += r.NetPay
		}
		fmt.Printf("Total Net Pay: %.2f\n", totalNetPay)
	case <-timeout:
		fmt.Println("Processing timed out")
	}
}
