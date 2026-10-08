package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var completedJobs int
var mutex sync.Mutex

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker", id, "stopped")
			return

		case job, ok := <-jobs:
			if !ok {
				fmt.Println("Worker", id, "finished")
				return
			}

			fmt.Println("Worker", id, "processing job", job)

			time.Sleep(500 * time.Millisecond)

			mutex.Lock()
			completedJobs++
			mutex.Unlock()

			fmt.Println("Worker", id, "completed job", job)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	jobs := make(chan int)

	var wg sync.WaitGroup

	wg.Add(3)

	go worker(ctx, 1, jobs, &wg)
	go worker(ctx, 2, jobs, &wg)
	go worker(ctx, 3, jobs, &wg)

	for i := 1; i <= 10; i++ {
		select {
		case jobs <- i:
			fmt.Println("Sent job", i)

		case <-ctx.Done():
			fmt.Println("Job sending stopped")
			close(jobs)
			wg.Wait()
			fmt.Println("Completed jobs:", completedJobs)
			return
		}
	}

	close(jobs)

	wg.Wait()

	fmt.Println("Completed jobs:", completedJobs)
	fmt.Println("Program completed")
}
