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
}

type ProcessedResult struct {
	WorkerID    int
	Employee    Employee
	FinalPayout float64
}

func employeeProducer(jobs chan<- Employee, totalEmployees int) {
	fmt.Println("[Producer] Starting to generate employee records...")
	
	names := []string{"Alice", "Bob", "Charlie", "Diana", "Ethan", "Fiona", "George"}
	
	for i := 1; i <= totalEmployees; i++ {
		emp := Employee{
			ID:         100 + i,
			Name:       names[(i-1)%len(names)],
			BaseSalary: 50000 + float64(rand.Intn(40000)),
		}
		jobs <- emp
		time.Sleep(100 * time.Millisecond) // Simulate generation pacing
	}
	
	close(jobs)
	fmt.Println("[Producer] Finished sending all employee records.")
}

func employeeWorker(id int, jobs <-chan Employee, results chan<- ProcessedResult, alerts chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for emp := range jobs {
		// Simulate processing duration
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)

		// Business rule check: Flag salary anomalies as alerts
		if emp.BaseSalary > 85000 {
			alerts <- fmt.Sprintf("Worker %d flagged high salary anomaly for %s (ID: %d, Salary: $%.2f)", id, emp.Name, emp.ID, emp.BaseSalary)
			continue
		}

		// Standard calculation (Bonus and Tax)
		bonus := emp.BaseSalary * 0.12
		tax := emp.BaseSalary * 0.18
		netPayout := emp.BaseSalary + bonus - tax

		results <- ProcessedResult{
			WorkerID:    id,
			Employee:    emp,
			FinalPayout: netPayout,
		}
	}
}

func main() {
	const totalEmployees = 7
	const numWorkers = 3

	jobs := make(chan Employee, totalEmployees)
	results := make(chan ProcessedResult, totalEmployees)
	alerts := make(chan string, totalEmployees)

	var wg sync.WaitGroup

	// 1. Boot up Worker Pool
	fmt.Printf("Booting up %d employee processing workers...\n\n", numWorkers)
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go employeeWorker(w, jobs, results, alerts, &wg)
	}

	// 2. Start Producer
	go employeeProducer(jobs, totalEmployees)

	// Background closer for results and alerts once workers finish
	go func() {
		wg.Wait()
		close(results)
		close(alerts)
	}()

	// 3. Consumer / Aggregator loop using Select with a timeout mechanism
	// We expect a total of 'totalEmployees' outputs across results and alerts combined.
	for i := 0; i < totalEmployees; i++ {
		select {
		case res, ok := <-results:
			if ok {
				fmt.Printf(" SUCCESS: Worker %d processed %s (ID: %d) | Net Payout: $%.2f\n",
					res.WorkerID, res.Employee.Name, res.Employee.ID, res.FinalPayout)
			}
		case alert, ok := <-alerts:
			if ok {
				fmt.Printf(" ALERT:   %s\n", alert)
			}
		case <-time.After(3 * time.Second):
			fmt.Println(" TIMEOUT: System processing is running too slow, aborting wait.")
			return
		}
	}

	fmt.Println("\nSystem shutdown complete. All employees processed.")
}