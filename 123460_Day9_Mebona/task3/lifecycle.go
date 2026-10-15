package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		// 1. WAITING / SLEEPING phase
		fmt.Println("Goroutine: Waiting...")
		time.Sleep(100 * time.Millisecond)

		// 2. EXECUTION phase
		fmt.Println("Goroutine: Executing work...")

		// 3. TERMINATION phase
		fmt.Println("Goroutine: Terminating.")
		wg.Done()
	}()

	fmt.Println("Main: Waiting for goroutine to finish...")
	wg.Wait()
	fmt.Println("Main: Program finished.")
}