package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// worker processes jobs.
// Multiple workers form the worker pool (fan-out).
func worker(
	ctx context.Context,
	id int,
	jobs <-chan int,
	results chan<- int,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for {
		select {
		// Stop when context is cancelled or times out.
		case <-ctx.Done():
			fmt.Println("Worker", id, "cancelled")
			return

		// Receive a job.
		case job, ok := <-jobs:
			if !ok {
				fmt.Println("Worker", id, "finished")
				return
			}

			fmt.Println("Worker", id, "processing job", job)

			// Simulate work.
			time.Sleep(500 * time.Millisecond)

			result := job * 2

			// Send result unless context is cancelled.
			select {
			case results <- result:
				fmt.Println("Worker", id, "sent result", result)

			case <-ctx.Done():
				fmt.Println("Worker", id, "stopped")
				return
			}
		}
	}
}

func main() {
	// Timeout after 3 seconds.
	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)
	defer cancel()

	// Jobs channel.
	jobs := make(chan int)

	// Results channel.
	results := make(chan int)

	var wg sync.WaitGroup

	// Worker pool: 3 workers.
	wg.Add(3)

	go worker(ctx, 1, jobs, results, &wg)
	go worker(ctx, 2, jobs, results, &wg)
	go worker(ctx, 3, jobs, results, &wg)

	// FAN-OUT:
	// Send jobs to the shared jobs channel.
	go func() {
		defer close(jobs)

		for i := 1; i <= 10; i++ {
			select {
			case jobs <- i:
				fmt.Println("Job sent:", i)

			case <-ctx.Done():
				fmt.Println("Job producer cancelled")
				return
			}
		}
	}()

	// Close results after every worker finishes.
	go func() {
		wg.Wait()
		close(results)
	}()

	// FAN-IN:
	// Collect results from all workers.
	for result := range results {
		fmt.Println("Final result:", result)
	}

	fmt.Println("Program finished")
}

// Race-safe processing:
// Workers communicate through channels instead of directly modifying
// shared variables. Run with:
//
// go run -race main.go
//
// This project demonstrates:
// 1. Worker pool
// 2. Fan-out
// 3. Fan-in
// 4. Context timeout/cancellation
// 5. Race-safe processing
