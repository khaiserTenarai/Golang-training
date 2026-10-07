package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID    int
	Value int
}

type Result struct {
	TaskID int
	Value  int
}

type Stats struct {
	mu        sync.Mutex
	Processed int
}

func (s *Stats) Inc() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Processed++
}

func generator(ctx context.Context, nums ...int) <-chan Task {
	out := make(chan Task)
	go func() {
		defer close(out)
		for i, n := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- Task{ID: i+1, Value: n}:
			}
		}
	}()
	return out
}

func worker(ctx context.Context, id int, jobs <-chan Task, stats *Stats) <-chan Result {
	out := make(chan Result)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case job, ok := <-jobs:
				if !ok {
					return
				}
				time.Sleep(100 * time.Millisecond)
				stats.Inc()
				
				select {
				case <-ctx.Done():
					return
				case out <- Result{TaskID: job.ID, Value: job.Value * 2}:
				}
			}
		}
	}()
	return out
}

func fanIn(ctx context.Context, channels ...<-chan Result) <-chan Result {
	var wg sync.WaitGroup
	out := make(chan Result)

	multiplex := func(c <-chan Result) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case res, ok := <-c:
				if !ok {
					return
				}
				select {
				case <-ctx.Done():
					return
				case out <- res:
				}
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer cancel()

	stats := &Stats{}

	jobs := generator(ctx, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	numWorkers := 3
	workerChannels := make([]<-chan Result, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workerChannels[i] = worker(ctx, i+1, jobs, stats)
	}

	results := fanIn(ctx, workerChannels...)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Main: Timeout reached, shutting down")
			fmt.Printf("Total processed safely: %d\n", stats.Processed)
			return
		case res, ok := <-results:
			if !ok {
				fmt.Println("Main: All jobs completed")
				fmt.Printf("Total processed safely: %d\n", stats.Processed)
				return
			}
			fmt.Printf("Processed Task ID: %d, Result: %d\n", res.TaskID, res.Value)
		}
	}
}