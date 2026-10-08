package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(id int, jobs chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		jobID := id*10 + i
		fmt.Printf("[Producer %d] Produced job %d\n", id, jobID)
		jobs <- jobID
		time.Sleep(100 * time.Millisecond)
	}
}

func consumer(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("   [Consumer %d] Processing job %d...\n", id, job)
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("   [Consumer %d] Completed job %d\n", id, job)
	}
}

func main() {

	jobs := make(chan int, 5)

	var producerWG sync.WaitGroup
	var consumerWG sync.WaitGroup

	for p := 1; p <= 2; p++ {
		producerWG.Add(1)
		go producer(p, jobs, &producerWG)
	}

	for c := 1; c <= 2; c++ {
		consumerWG.Add(1)
		go consumer(c, jobs, &consumerWG)
	}

	go func() {
		producerWG.Wait()
		close(jobs)
	}()

	consumerWG.Wait()

	fmt.Println("[Main] All production and consumption completed successfully.")
}