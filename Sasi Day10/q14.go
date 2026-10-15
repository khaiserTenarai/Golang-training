package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count int64
	var wg sync.WaitGroup

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			
			// RACEY CODE (triggers go run -race):
			// count++

			// SAFE FIX (Race-free):
			atomic.AddInt64(&count, 1)
		}()
	}

	wg.Wait()
	fmt.Println("Final Count:", atomic.LoadInt64(&count))
}