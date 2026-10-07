package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("Worker", id, "processing job", job)
		result := job * 2
		results <- result
	}
}

func main() {
	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup

	// Create 3 workers.
	wg.Add(3)

	go worker(1, jobs, results, &wg)
	go worker(2, jobs, results, &wg)
	go worker(3, jobs, results, &wg)

	// Send jobs.
	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// Close results after all workers finish.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Read results.
	for result := range results {
		fmt.Println("Result:", result)
	}
}
