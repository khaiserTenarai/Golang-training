package main

import (
	"fmt"
	"sync"
)
// The mutex helps prevent a race condition.
func main() {
	/*
		Problem:
		Multiple goroutines are changing the same variable.

		This can cause a race condition.

		Solution:
		Use sync.Mutex to allow only one goroutine
		to change the shared data at a time.

		Lock()   -> locks the shared data
		Unlock() -> unlocks the shared data
	*/

	var mutex sync.Mutex
	var count int
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mutex.Lock()
			count++
			mutex.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println("Final count:", count)
}