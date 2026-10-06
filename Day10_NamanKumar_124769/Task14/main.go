package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++ // DATA RACE
		}()
	}
	wg.Wait()
	fmt.Println(count) // unpredictable
}
