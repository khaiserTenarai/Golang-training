package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Job represents one job.
type Job struct {
	ID int
}

// Result represents the result of a job.
type Result struct {
	JobID int
	Value int
}

/*
	Worker processes jobs.

	Multiple workers receive jobs from the same channel.
	This is the FAN-OUT part.

	Context allows workers to stop when timeout/cancellation occurs.
*/
func worker(
	ctx context.Context,
	id int,
	jobs <-chan Job,
	results chan<- Result,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for {
		select {

		case <-ctx.Done():
			fmt.Println("Worker", id, "stopped:", ctx.Err())
			return

		case job, ok := <-jobs:
			if !ok {
				return
			}

			// Simulate some work.
			time.Sleep(500 * time.Millisecond)

			/*
				Each worker creates its own result.

				There is no shared counter here,
				so this part is race-safe.
			*/
			result := Result{
				JobID: job.ID,
				Value: job.ID * 10,
			}

			select {
			case results <- result:
				fmt.Println(
					"Worker", id,
					"processed job", job.ID,
				)

			case <-ctx.Done():
				fmt.Println("Worker", id, "cancelled")
				return
			}
		}
	}
}

func main() {
	/*
		DAY 10 MINI PROJECT

		Requirements:

		1. Worker Pool
		2. Fan-Out
		3. Fan-In
		4. Timeout
		5. Race-safe processing

		Flow:

		Jobs
		  ↓
		Worker Pool
		  ↓
		Multiple Workers
		  ↓
		Results Channel
		  ↓
		Main

		Context controls cancellation.
	*/

	// Create a context with 3-second timeout.
	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)

	defer cancel()

	// Jobs channel.
	jobs := make(chan Job)

	// Results channel.
	results := make(chan Result)

	var wg sync.WaitGroup

	// -----------------------------------
	// FAN-OUT: Create multiple workers
	// -----------------------------------

	workerCount := 3

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)

		go worker(
			ctx,
			i,
			jobs,
			results,
			&wg,
		)
	}

	// -----------------------------------
	// Send jobs
	// -----------------------------------

	go func() {
		defer close(jobs)

		for i := 1; i <= 10; i++ {

			select {
			case jobs <- Job{ID: i}:

				fmt.Println("Job sent:", i)

			case <-ctx.Done():

				fmt.Println("Job producer stopped")
				return
			}
		}
	}()

	// -----------------------------------
	// FAN-IN: Collect results
	// -----------------------------------

	go func() {
		wg.Wait()

		// All workers finished.
		close(results)
	}()

	// -----------------------------------
	// Read final results
	// -----------------------------------

	for result := range results {
		fmt.Println(
			"Result received:",
			result.JobID,
			"->",
			result.Value,
		)
	}

	fmt.Println("Day 10 project completed")
}