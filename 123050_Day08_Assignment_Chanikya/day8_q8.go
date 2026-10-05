package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 100; i++ {
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
