package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Job struct{ ID int }

type Result struct {
	JobID    int
	WorkerID int
	Value    int
}

// Race-safe stats
type Stats struct {
	processed atomic.Int64
	mu        sync.Mutex
	perWorker map[int]int
}

func (s *Stats) Record(workerID int) {
	s.processed.Add(1)
	s.mu.Lock()
	s.perWorker[workerID]++
	s.mu.Unlock()
}

// Producer: sends jobs until done or cancelled
func produce(ctx context.Context, n int) <-chan Job {
	jobs := make(chan Job)
	go func() {
		defer close(jobs)
		for i := 1; i <= n; i++ {
			select {
			case jobs <- Job{ID: i}:
			case <-ctx.Done():
				return
			}
		}
	}()
	return jobs
}

// One worker: reads from shared jobs channel (fan-out), returns its own result channel
func worker(ctx context.Context, id int, jobs <-chan Job, stats *Stats) <-chan Result {
	out := make(chan Result)
	go func() {
		defer close(out)
		for job := range jobs {
			// simulate work, but stay cancellable
			select {
			case <-time.After(time.Duration(100+rand.Intn(300)) * time.Millisecond):
			case <-ctx.Done():
				return
			}
			res := Result{JobID: job.ID, WorkerID: id, Value: job.ID * job.ID}
			select {
			case out <- res:
				stats.Record(id)
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

// Fan-in: merge all worker outputs
func merge(ctx context.Context, chans ...<-chan Result) <-chan Result {
	out := make(chan Result)
	var wg sync.WaitGroup
	for _, c := range chans {
		wg.Add(1)
		go func(c <-chan Result) {
			defer wg.Done()
			for r := range c {
				select {
				case out <- r:
				case <-ctx.Done():
					return
				}
			}
		}(c)
	}
	go func() { wg.Wait(); close(out) }()
	return out
}

func main() {
	const (
		numJobs    = 20
		numWorkers = 4
		timeout    = 1 * time.Second
	)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	stats := &Stats{perWorker: make(map[int]int)}

	jobs := produce(ctx, numJobs)

	// Fan-out
	workers := make([]<-chan Result, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workers[i] = worker(ctx, i+1, jobs, stats)
	}

	// Fan-in
	for r := range merge(ctx, workers...) {
		fmt.Printf("job %2d -> %3d (worker %d)\n", r.JobID, r.Value, r.WorkerID)
	}

	if ctx.Err() != nil {
		fmt.Println("stopped early:", ctx.Err())
	} else {
		fmt.Println("all jobs completed")
	}

	stats.mu.Lock()
	defer stats.mu.Unlock()
	fmt.Println("processed:", stats.processed.Load(), "per worker:", stats.perWorker)
}
