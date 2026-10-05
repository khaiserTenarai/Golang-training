package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Worker", id, "started")
	time.Sleep(1 * time.Second)
	fmt.Println("Worker", id, "finished")
}

func main() {

	var wg sync.WaitGroup

	// Create 3 goroutines
	wg.Add(3)

	go worker(1, &wg)
	go worker(2, &wg)
	go worker(3, &wg)

	// Wait until all goroutines finish
	wg.Wait()

	fmt.Println("All workers completed")
}