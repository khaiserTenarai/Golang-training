package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter++ // Race condition
			}
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
