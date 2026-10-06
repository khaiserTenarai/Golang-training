package main
// A worker pool means we have a fixed number of workers processing multiple jobs.
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
	/*
		Problem:
		We have many jobs to process.

		Solution:
		Create a fixed number of workers.

		Here:
		3 workers process 6 jobs.
	*/

	jobs := make(chan int, 6)

	var wg sync.WaitGroup

	// Start 3 workers.
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Send 6 jobs.
	for i := 1; i <= 6; i++ {
		jobs <- i
	}

	// Close jobs channel.
	close(jobs)

	// Wait for workers.
	wg.Wait()

	fmt.Println("All jobs completed")
}