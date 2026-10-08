package main

import (
	"fmt"
	"sync"
)

func main() {
	work := make(chan int, 10)
	var wg sync.WaitGroup

	// Populate work channel
	for i := 1; i <= 5; i++ {
		work <- i
	}
	close(work)

	// Fan-out: Spawn multiple goroutines reading from the same channel
	numWorkers := 3
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for item := range work {
				fmt.Printf("Worker %d processed task %d\n", workerID, item)
			}
		}(w)
	}

	wg.Wait()
}