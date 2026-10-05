package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Result struct {
	WorkerID int
	JobID    int
	Message  string
}

func worker(
	ctx context.Context,
	id int,
	jobs <-chan int,
	results chan<- Result,
	wg *sync.WaitGroup,
	counter *int64,
) {
	defer wg.Done()

	for {

		select {

		case <-ctx.Done():
			fmt.Println("Worker", id, "cancelled")
			return

		case job, ok := <-jobs:

			if !ok {
				fmt.Println("Worker", id, "finished")
				return
			}

			fmt.Println("Worker", id, "processing job", job)

			// Simulate work
			time.Sleep(500 * time.Millisecond)

			// Race-safe counter
			atomic.AddInt64(counter, 1)

			result := Result{
				WorkerID: id,
				JobID:    job,
				Message:  "Job completed",
			}

			select {

			case results <- result:

			case <-ctx.Done():
				fmt.Println("Worker", id, "stopped while sending result")
				return
			}
		}
	}
}

func main() {

	// Create timeout context
	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)

	defer cancel()

	jobs := make(chan int)
	results := make(chan Result)

	var wg sync.WaitGroup

	var completedJobs int64

	// Worker pool
	numberOfWorkers := 3

	wg.Add(numberOfWorkers)

	for i := 1; i <= numberOfWorkers; i++ {
		go worker(
			ctx,
			i,
			jobs,
			results,
			&wg,
			&completedJobs,
		)
	}

	// Fan-out: send jobs to multiple workers
	go func() {

		defer close(jobs)

		for job := 1; job <= 10; job++ {

			select {

			case jobs <- job:

			case <-ctx.Done():
				fmt.Println("Job producer cancelled")
				return
			}
		}

	}()

	// Fan-in: collect results
	go func() {

		wg.Wait()

		close(results)

	}()

	// Read results
	for result := range results {

		fmt.Println(
			"Result:",
			"Worker", result.WorkerID,
			"Job", result.JobID,
			result.Message,
		)
	}

	fmt.Println("Completed Jobs:", completedJobs)

	fmt.Println("Mini project finished")
}