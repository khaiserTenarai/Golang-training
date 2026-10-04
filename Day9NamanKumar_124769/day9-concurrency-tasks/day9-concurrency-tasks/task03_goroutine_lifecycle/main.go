package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Goroutine %d: running\n", id)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Goroutine %d: completed\n", id)
}

func main() {
	fmt.Println("Goroutines before creation:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	time.Sleep(10 * time.Millisecond)
	fmt.Println("Goroutines while running:", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("Goroutines after completion:", runtime.NumGoroutine())
}
