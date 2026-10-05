package main

import (
	"fmt"
	"sync"
)

func main() {

	counter := 0

	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(5)

	for i := 1; i <= 5; i++ {

		go func() {
			defer wg.Done()

			// Protect shared counter
			mu.Lock()

			counter++

			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println("Final Counter:", counter)
}