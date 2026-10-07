package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)

		results <- job * 2
	}
}

func main() {
	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup

	// Start workers.
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs.
	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- i
		}

		close(jobs)
	}()

	// Close results after workers finish.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Read results.
	for result := range results {
		fmt.Println("Result:", result)
	}
}
