package main

import (
	"fmt"  // Imports fmt for output.
	"sync" // Imports sync for WaitGroup.
	"time" // Imports time for simulated work.
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) { // Defines a worker receiving jobs.
	defer wg.Done()         // Signals worker completion.
	for job := range jobs { // Reads jobs until the channel closes.
		fmt.Println("Worker", id, "processing job", job) // Prints job assignment.
		time.Sleep(100 * time.Millisecond)               // Simulates processing.
		fmt.Println("Worker", id, "completed job", job)  // Prints completion.
	}
}

func main() {
	jobs := make(chan int) // Creates the job channel.
	var wg sync.WaitGroup  // Creates a WaitGroup.
	workers := 3           // Defines the number of workers.

	wg.Add(workers) // Registers all workers.

	for i := 1; i <= workers; i++ { // Creates workers.
		go worker(i, jobs, &wg) // Starts each worker.
	}

	for job := 1; job <= 10; job++ { // Produces ten jobs.
		jobs <- job // Sends a job to an available worker.
	}

	close(jobs)                       // Closes the channel after all jobs are sent.
	wg.Wait()                         // Waits for all workers.
	fmt.Println("All jobs completed") // Prints final status.
}

// example-dominos
