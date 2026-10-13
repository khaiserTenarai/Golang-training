package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processed job %d\n", id, job)
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

	// Send jobs to workers
	for i := 1; i <= 5; i++ {
		jobs <- i
	}

	close(jobs)
	wg.Wait()
}
