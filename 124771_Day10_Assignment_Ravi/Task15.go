package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID int
}

type Result struct {
	JobID  int
	Output string
}

func worker(
	ctx context.Context,
	id int,
	jobs <-chan Job,
	results chan<- Result,
	wg *sync.WaitGroup,
	mu *sync.Mutex,
	processed *int,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}

			select {
			case <-time.After(300 * time.Millisecond):
				mu.Lock()
				(*processed)++
				mu.Unlock()

				select {
				case results <- Result{
					JobID: job.ID,
					Output: fmt.Sprintf(
						"Job %d processed by worker %d",
						job.ID, id,
					),
				}:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(), 2*time.Second,
	)
	defer cancel()

	jobs := make(chan Job)
	results := make(chan Result)

	const workerCount = 3
	var workerWG sync.WaitGroup
	var resultWG sync.WaitGroup
	var mu sync.Mutex
	processed := 0

	// Fan-out: start the worker pool.
	for i := 1; i <= workerCount; i++ {
		workerWG.Add(1)
		go worker(ctx, i, jobs, results, &workerWG, &mu, &processed)
	}

	// Submit jobs until completion or cancellation.
	go func() {
		defer close(jobs)

		for i := 1; i <= 10; i++ {
			select {
			case jobs <- Job{ID: i}:
			case <-ctx.Done():
				return
			}
		}
	}()

	// Close results after all workers finish.
	go func() {
		workerWG.Wait()
		close(results)
	}()

	// Fan-in: collect results from all workers.
	resultWG.Add(1)
	go func() {
		defer resultWG.Done()

		for result := range results {
			fmt.Println(result.Output)
		}
	}()

	resultWG.Wait()

	mu.Lock()
	fmt.Println("Total jobs processed:", processed)
	mu.Unlock()

	if ctx.Err() != nil {
		fmt.Println("Processing stopped:", ctx.Err())
	} else {
		fmt.Println("All jobs completed")
	}
}
