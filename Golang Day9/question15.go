package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


type Employee struct {
	ID   int
	Name string
}


func producer(jobs chan<- Employee, numJobs int) {
	fmt.Println("[Producer] Starting to generate employee records...")
	for i := 1; i <= numJobs; i++ {
		jobs <- Employee{ID: 100 + i, Name: fmt.Sprintf("Employee-%d", i)}
	}
	
	close(jobs)
	fmt.Println("[Producer] Finished sending all records.")
}

func worker(id int, jobs <-chan Employee, results chan<- string, alerts chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for emp := range jobs {
		
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

		
		if emp.ID%3 == 0 {
			alerts <- fmt.Sprintf("Worker %d flagged invalid data for %s (ID: %d)", id, emp.Name, emp.ID)
		} else {
			results <- fmt.Sprintf("Worker %d successfully processed %s (ID: %d)", id, emp.Name, emp.ID)
		}
	}
}

func main() {
	const totalEmployees = 7
	const numWorkers = 3

	
	jobs := make(chan Employee, totalEmployees)
	results := make(chan string, totalEmployees)
	alerts := make(chan string, totalEmployees)

	var wg sync.WaitGroup

	
	fmt.Printf("Booting up %d workers...\n\n", numWorkers)
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, alerts, &wg)
	}

	
	go producer(jobs, totalEmployees)

	
	for i := 0; i < totalEmployees; i++ {
		select {
		case res := <-results:
			fmt.Println(" SUCCESS:", res)
		case alert := <-alerts:
			fmt.Println("ALERT:  ", alert)
		case <-time.After(2 * time.Second):
			
			fmt.Println("TIMEOUT: System is running too slow, aborting wait.")
		}
	}

	wg.Wait()
	close(results)
	close(alerts)
	
	fmt.Println("\nSystem shutdown complete. All employees accounted for.")
}