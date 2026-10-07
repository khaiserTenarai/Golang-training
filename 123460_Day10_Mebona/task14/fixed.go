package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++ // Protected by Mutex
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Count:", count)
}