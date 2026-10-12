//Fan-out means one source of work sends tasks to multiple goroutines
// so they can process the work at the same time.

package main

import (
	"fmt"  // Imports fmt for printing output.
	"sync" // Imports sync for WaitGroup.
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {

	defer wg.Done() // Tells WaitGroup that this worker is finished.

	for job := range jobs { // Receives jobs from the channel.

		fmt.Println("Worker", id, "processed job", job) // Prints the worker and job.
	}
}

func main() {

	jobs := make(chan int) // Creates a channel for jobs.

	var wg sync.WaitGroup // Creates a WaitGroup.

	wg.Add(3) // Registers three worker goroutines.

	// Starts three workers.
	go worker(1, jobs, &wg)
	go worker(2, jobs, &wg)
	go worker(3, jobs, &wg)

	// Sends jobs to the workers.
	for i := 1; i <= 6; i++ {

		jobs <- i // Sends one job to the channel.
	}

	close(jobs) // Closes the channel after all jobs are sent.

	wg.Wait() // Waits for all workers to finish.

	fmt.Println("All jobs completed") // Prints when all workers are done.
}
