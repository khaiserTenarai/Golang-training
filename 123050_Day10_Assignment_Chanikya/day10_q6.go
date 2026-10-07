package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter atomic.Int64

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			counter.Add(1)
		}()
	}

	wg.Wait()

	fmt.Println("Counter:", counter.Load())
}
