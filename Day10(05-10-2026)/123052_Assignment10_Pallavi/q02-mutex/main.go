package main

import (
	"fmt"
	"sync"
)

func main() {

	var mu sync.Mutex
	counter := 0

	var wg sync.WaitGroup

	wg.Add(5)

	for i := 1; i <= 5; i++ {
		go func() {
			defer wg.Done()

			// Lock before modifying shared data
			mu.Lock()

			counter++

			// Unlock after modifying shared data
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println("Final Counter:", counter)
}