package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

func increase(wg *sync.WaitGroup) {
	defer wg.Done()

	// Atomically increase counter by 1.
	atomic.AddInt64(&counter, 1)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go increase(&wg)
	}

	wg.Wait()

	fmt.Println("Final counter:", counter)
}
