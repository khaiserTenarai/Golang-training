package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // DATA RACE: unsynchronized write
		}()
	}
	wg.Wait()
	fmt.Println("counter:", counter)
}