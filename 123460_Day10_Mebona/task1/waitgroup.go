package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement counter when worker finishes
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment counter for each goroutine
		go worker(i, &wg)
	}

	wg.Wait() // Block until counter is zero
	fmt.Println("All workers completed.")
}