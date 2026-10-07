package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func increase(wg *sync.WaitGroup) {
	defer wg.Done()

	mutex.Lock() // Only one goroutine can enter the critical section.
	counter++
	mutex.Unlock() // Allow another goroutine to enter.
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
