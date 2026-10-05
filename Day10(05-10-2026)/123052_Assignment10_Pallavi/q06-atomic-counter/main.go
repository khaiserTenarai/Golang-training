package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var counter int64

	var wg sync.WaitGroup

	wg.Add(5)

	for i := 1; i <= 5; i++ {

		go func() {
			defer wg.Done()

			// Safely increment shared counter
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()

	fmt.Println("Final Counter:", counter)
}