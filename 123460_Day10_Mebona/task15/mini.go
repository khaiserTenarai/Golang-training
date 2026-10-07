package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Task represents a unit of work with a specific ID.
type Task struct {
	ID int
}

// Result represents the outcome of a processed task.
type Result struct {
	TaskID   int
	WorkerID int
	Value    int
	Err      error
}

// Stats safely tracks processed metrics across concurrent workers using atomics.
type Stats struct {
	successCount int64
	failureCount int64
}

func (s *Stats) IncSuccess() { atomic.AddInt64(&s.successCount, 1) }
func (s *Stats) IncFailure() { atomic.AddInt64(&s.failureCount, 1) }
func (s *Stats) GetSuccess() int64 { return atomic.LoadInt64(&s.successCount) }
func (s *Stats) GetFailure() int64 { return atomic.LoadInt64(&s.failureCount) }

// worker represents an individual processor in the Worker Pool (Fan-Out target).
func worker(ctx context.Context, id int, jobs <-chan Task, results chan<- Result, stats *Stats, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			// Context has been cancelled or timed out
			return
		case task, ok := <-jobs:
			if !ok {
				// Channel closed, terminate worker
				return
			}

			// Simulate processing time or cooperative cancellation check
			select {
			case <-time.After(150 * time.Millisecond):
				// Task successfully completed
				res := Result{
					TaskID:   task.ID,
					WorkerID: id,
					Value:    task.ID * 10,
				}
				stats.IncSuccess()
				results <- res
			case <-ctx.Done():
				// Interrupted mid-processing by timeout or cancellation
				stats.IncFailure()
				return
			}
		}
	}
}

// fanIn collects results from multiple worker channels into a single unified output stream.
func fanIn(channels ...<-chan Result) <-chan Result {
	out := make(chan Result)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan Result) {
			defer wg.Done()
			for res := range c {
				out <- res
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
	// Set a timeout of 600ms for the entire worker pipeline system
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Millisecond)
	defer cancel()

	numTasks := 10
	numWorkers := 3

	jobs := make(chan Task, numTasks)
	workerResults := make(chan Result, numTasks)
	stats := &Stats{}

	// Populate jobs (Fan-Out distribution source)
	for i := 1; i <= numTasks; i++ {
		jobs <- Task{ID: i}
	}
	close(jobs)

	// Launch Worker Pool (Fan-Out pattern: multiple workers listening on the same 'jobs' channel)
	var workerWg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		workerWg.Add(1)
		go worker(ctx, w, jobs, workerResults, stats, &workerWg)
	}

	// Close workerResults channel once all workers finish executing
	go func() {
		workerWg.Wait()
		close(workerResults)
	}()

	// Fan-In: Consolidate results stream
	finalOutput := fanIn(workerResults)

	// Read and print incoming results from the unified stream
	var collectorWg sync.WaitGroup
	collectorWg.Add(1)
	go func() {
		defer collectorWg.Done()
		for res := range finalOutput {
			if res.Err != nil {
				fmt.Printf("[Error] Task %d failed\n", res.TaskID)
			} else {
				fmt.Printf("[Success] Task %d processed by Worker %d -> Result: %d\n", res.TaskID, res.WorkerID, res.Value)
			}
		}
	}()

	collectorWg.Wait()
	fmt.Printf("Execution Complete.\nTotal Successes (Atomic): %d\nTotal Failures/Cancelled (Atomic): %d\n", stats.GetSuccess(), stats.GetFailure())
}