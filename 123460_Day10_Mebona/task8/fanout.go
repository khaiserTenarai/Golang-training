package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int, 5)

	// Populate jobs
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	var wg sync.WaitGroup

	// Fan-out: Spin up multiple goroutines reading from the SAME channel
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				fmt.Printf("Worker %d executed job %d\n", workerID, job)
			}
		}(w)
	}

	wg.Wait()
}