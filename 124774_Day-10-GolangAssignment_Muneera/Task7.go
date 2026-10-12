package main

import (
	"fmt"  // Imports fmt for output.
	"sync" // Imports sync for WaitGroup.
)

// worker processes jobs received from the jobs channel.
func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {

	defer wg.Done() // Signals that this worker is finished.

	for job := range jobs { // Receives jobs until the channel is closed.

		fmt.Println("Worker", id, "processing job", job) // Prints which worker is processing the job.
	}
}

func main() {

	jobs := make(chan int) // Creates a channel for jobs.

	var wg sync.WaitGroup // Creates a WaitGroup.

	wg.Add(3) // Registers three workers.

	// Starts three workers.
	for i := 1; i <= 3; i++ {

		go worker(i, jobs, &wg) // Starts a worker goroutine.
	}

	// Sends jobs to the workers.
	go func() {

		for i := 1; i <= 6; i++ { // Creates six jobs.

			jobs <- i // Sends the job to the workers.
		}

		close(jobs) // Closes the jobs channel after all jobs are sent.
	}()

	wg.Wait() // Waits for all workers to finish.

	fmt.Println("All jobs completed") // Prints when all jobs are finished.
}
