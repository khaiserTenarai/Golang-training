package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Job struct {
	ID    int
	Value int
}

type Result struct {
	JobID  int
	Output int
	Err    error
}

type System struct {
	processedJobs int64
}

func (s *System) worker(ctx context.Context, id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}

			// Simulate workload with variable latency
			processingTime := time.Duration(50+rand.Intn(100)) * time.Millisecond

			select {
			case <-time.After(processingTime):
				atomic.AddInt64(&s.processedJobs, 1)
				results <- Result{
					JobID:  job.ID,
					Output: job.Value * 10,
				}
			case <-ctx.Done():
				results <- Result{
					JobID: job.ID,
					Err:   ctx.Err(),
				}
				return
			}
		}
	}
}

// Fan-In: Merges multiple result channels into one
func fanInResults(ctx context.Context, channels ...<-chan Result) <-chan Result {
	out := make(chan Result)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan Result) {
			defer wg.Done()
			for res := range c {
				select {
				case <-ctx.Done():
					return
				case out <- res:
				}
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	const numWorkers = 3
	const numJobs = 10
	const timeout = 300 * time.Millisecond

	sys := &System{}

	// System-wide context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	jobsCh := make(chan Job, numJobs)
	
	// Create worker channels for fan-out
	workerChs := make([]chan Result, numWorkers)
	for i := range workerChs {
		workerChs[i] = make(chan Result, numJobs)
	}

	var wg sync.WaitGroup

	// Fan-out: Spin up worker pool
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go sys.worker(ctx, i+1, jobsCh, workerChs[i], &wg)
	}

	// Enqueue jobs
	for j := 1; j <= numJobs; j++ {
		jobsCh <- Job{ID: j, Value: j}
	}
	close(jobsCh)

	// Close worker result channels when workers finish
	go func() {
		wg.Wait()
		for _, ch := range workerChs {
			close(ch)
		}
	}()

	// Convert []chan Result into []<-chan Result for Fan-In
	readOnlyChs := make([]<-chan Result, numWorkers)
	for i, ch := range workerChs {
		readOnlyChs[i] = ch
	}

	// Fan-in: Consume consolidated results
	mergedResults := fanInResults(ctx, readOnlyChs...)

	fmt.Println("--- Processing Pipeline Output ---")
	for res := range mergedResults {
		if res.Err != nil {
			fmt.Printf("Job %2d: CANCELLED (%v)\n", res.JobID, res.Err)
		} else {
			fmt.Printf("Job %2d: SUCCESS   Output = %d\n", res.JobID, res.Output)
		}
	}

	fmt.Printf("\nTotal Jobs Successfully Processed (Race-Safe): %d\n", atomic.LoadInt64(&sys.processedJobs))
}