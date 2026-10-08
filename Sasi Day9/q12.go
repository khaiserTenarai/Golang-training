package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("   [Worker %d] Processing job %d...\n", id, job)
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("   [Worker %d] Finished job %d\n", id, job)
	}
}

func main() {

	const bufferCapacity = 3
	jobs := make(chan int, bufferCapacity)

	var wg sync.WaitGroup

	for w := 1; w <= 2; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	totalJobs := 8
	for i := 1; i <= totalJobs; i++ {
		fmt.Printf("[Producer] Submitting job %d (Buffer count: %d/%d)\n", i, len(jobs), bufferCapacity)

		jobs <- i

		fmt.Printf("[Producer] Job %d accepted\n", i)
	}

	close(jobs)
	wg.Wait()

	fmt.Println("[Main] All jobs processed safely with backpressure control.")
}