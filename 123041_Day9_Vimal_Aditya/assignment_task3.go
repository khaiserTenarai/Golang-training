package main

import (
	"fmt"
	"sync"
	"time"
)

func workerLifecycle(id int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("[Worker %d] State: Started / Running\n", id)

	time.Sleep(300 * time.Millisecond)

	ch <- fmt.Sprintf("Worker %d completed its task successfully", id)

	fmt.Printf("[Worker %d] State: Finishing / Terminating\n", id)
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string, 2)

	fmt.Println("GOROUTINE LIFE CYCLE:")

	fmt.Println("[Main] Spawning Goroutine 1...")
	wg.Add(1)
	go workerLifecycle(1, ch, &wg)

	fmt.Println("[Main] Spawning Goroutine 2...")
	wg.Add(1)
	go workerLifecycle(2, ch, &wg)

	wg.Wait()
	close(ch)

	for msg := range ch {
		fmt.Println("[Main Received]:", msg)
	}

	fmt.Println("=== ALL GOROUTINES TERMINATED SAFELY ===")
}