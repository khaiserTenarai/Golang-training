package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("Worker", id, "processing", job)
	}
}

func main() {
	jobs := make(chan int)

	var wg sync.WaitGroup

	// Start multiple workers.
	wg.Add(3)

	go worker(1, jobs, &wg)
	go worker(2, jobs, &wg)
	go worker(3, jobs, &wg)

	// One source sends jobs to multiple workers.
	for i := 1; i <= 6; i++ {
		jobs <- i
	}

	close(jobs)

	wg.Wait()

	fmt.Println("All jobs completed")
}
