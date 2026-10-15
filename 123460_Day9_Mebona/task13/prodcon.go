package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(id int, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 2; i++ {
		out <- fmt.Sprintf("Job-%d-from-Producer-%d", i, id)
		time.Sleep(50 * time.Millisecond)
	}
}

func consumer(id int, in <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range in {
		fmt.Printf("Consumer %d handled: %s\n", id, job)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	jobs := make(chan string, 10)
	var prodWg, consWg sync.WaitGroup

	// Start producers
	for p := 1; p <= 2; p++ {
		prodWg.Add(1)
		go producer(p, jobs, &prodWg)
	}

	// Wait for producers to finish, then close channel
	go func() {
		prodWg.Wait()
		close(jobs)
	}()

	// Start consumers
	for c := 1; c <= 2; c++ {
		consWg.Add(1)
		go consumer(c, jobs, &consWg)
	}

	consWg.Wait()
	fmt.Println("All producer-consumer workflows finished.")
}