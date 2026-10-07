package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Result struct {
	WorkerID int
	JobID    int
	Value    int
}

type Processor struct {
	mu      sync.Mutex
	results []Result
}

func (p *Processor) Add(result Result) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.results = append(p.results, result)
}

func (p *Processor) Results() []Result {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Return a copy.
	results := make([]Result, len(p.results))
	copy(results, p.results)

	return results
}

func worker(
	ctx context.Context,
	id int,
	jobs <-chan int,
	results chan<- Result,
	processor *Processor,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d cancelled\n", id)
			return

		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d finished\n", id)
				return
			}

			fmt.Printf("Worker %d processing job %d\n", id, job)

			// Simulate work.
			select {
			case <-ctx.Done():
				fmt.Printf("Worker %d cancelled during job %d\n", id, job)
				return

			case <-time.After(500 * time.Millisecond):
			}

			result := Result{
				WorkerID: id,
				JobID:    job,
				Value:    job * job,
			}

			// Race-safe shared processing.
			processor.Add(result)

			// Fan-in.
			select {
			case results <- result:

			case <-ctx.Done():
				return
			}
		}
	}
}

func main() {
	// Timeout for the entire system.
	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)
	defer cancel()

	jobs := make(chan int)
	results := make(chan Result)

	processor := &Processor{}

	var wg sync.WaitGroup

	// Fan-out: start workers.
	workerCount := 3

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)

		go worker(
			ctx,
			i,
			jobs,
			results,
			processor,
			&wg,
		)
	}

	// Send jobs.
	go func() {
		defer close(jobs)

		for i := 1; i <= 20; i++ {
			select {
			case jobs <- i:
				fmt.Println("Submitted job:", i)

			case <-ctx.Done():
				fmt.Println("Job producer cancelled")
				return
			}
		}
	}()

	// Close results when all workers finish.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Fan-in: collect results.
	for result := range results {
		fmt.Printf(
			"Result: worker=%d job=%d value=%d\n",
			result.WorkerID,
			result.JobID,
			result.Value,
		)
	}

	fmt.Println("\nFinal processed results:")

	for _, result := range processor.Results() {
		fmt.Printf(
			"Worker %d processed Job %d => %d\n",
			result.WorkerID,
			result.JobID,
			result.Value,
		)
	}

	if ctx.Err() != nil {
		fmt.Println("\nSystem stopped:", ctx.Err())
	} else {
		fmt.Println("\nSystem completed successfully")
	}
}
