package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++ // Data race: Concurrent write without synchronization
		}()
	}

	wg.Wait()
	fmt.Println("Count:", count)
}