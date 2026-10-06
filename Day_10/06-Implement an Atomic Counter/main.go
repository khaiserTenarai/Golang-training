package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	/*
		Problem:
		Many goroutines are increasing the same counter.

		Solution:
		Use atomic.AddInt32().

		Atomic operation safely changes the value
		without using a Mutex.
	*/

	var counter int32
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// Safely increase counter by 1.
			atomic.AddInt32(&counter, 1)
		}()
	}

	wg.Wait()

	fmt.Println("Final counter:", counter)
}