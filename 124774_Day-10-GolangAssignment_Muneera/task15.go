package main

import (
	"context" // Imports context for cancellation and timeout.
	"fmt"     // Imports fmt for output.
	"sync"    // Imports sync for WaitGroup and Mutex.
	"time"    // Imports time for timeout and delay.
)

// Worker processes jobs.
func worker(ctx context.Context, id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done() // Marks the worker as finished.

	for {
		select {
		case <-ctx.Done(): // Stops worker when timeout happens.
			return

		case job, ok := <-jobs: // Receives a job.
			if !ok {
				return // Stops when jobs channel is closed.
			}

			mu.Lock()         // Locks shared processing.
			result := job * 2 // Processes the job safely.
			mu.Unlock()       // Releases the lock.

			results <- result // Sends result to fan-in channel.

			fmt.Println("Worker", id, "processed", job)
		}
	}
}

func main() {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second, // Sets a 2-second timeout.
	)
	defer cancel() // Releases context resources.

	jobs := make(chan int)    // Channel for jobs.
	results := make(chan int) // Channel for results.

	var wg sync.WaitGroup // Creates WaitGroup.
	var mu sync.Mutex     // Creates Mutex for race-safe processing.

	wg.Add(3) // Registers three workers.

	// Worker pool + fan-out.
	for i := 1; i <= 3; i++ {
		go worker(ctx, i, jobs, results, &wg, &mu) // Starts workers.
	}

	// Sends jobs to workers.
	go func() {
		defer close(jobs) // Closes jobs after sending.

		for i := 1; i <= 10; i++ {
			select {
			case jobs <- i: // Sends job to any available worker.
			case <-ctx.Done(): // Stops sending after timeout.
				return
			}
		}
	}()

	// Fan-in: waits for workers and closes results.
	go func() {
		wg.Wait()      // Waits for all workers.
		close(results) // Closes the results channel.
	}()

	// Reads combined results.
	for result := range results {
		fmt.Println("Result:", result) // Prints final result.
	}

	fmt.Println("Project finished")
}
