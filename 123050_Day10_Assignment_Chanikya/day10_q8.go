package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int)

	var wg sync.WaitGroup

	for workerID := 1; workerID <= 3; workerID++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			for job := range jobs {
				fmt.Printf("Worker %d processed %d\n", id, job)
			}
		}(workerID)
	}

	for i := 1; i <= 10; i++ {
		jobs <- i
	}

	close(jobs)

	wg.Wait()
}
