package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int64

func incrementCount(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		atomic.AddInt64(&count, 1)
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go increment(&wg)
	go increment(&wg)
	wg.Wait()
	fmt.Println("Final Counter: ", count)
}
