package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("Worker", id, "processing job", job)
	}
}

func main() {

	jobs := make(chan int)

	var wg sync.WaitGroup

	// Create 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Send jobs
	for job := 1; job <= 6; job++ {
		jobs <- job
	}

	// Close jobs channel
	close(jobs)

	// Wait for workers
	wg.Wait()

	fmt.Println("All jobs completed")
}