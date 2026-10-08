package main

import (
	"fmt"
	"sync"
)

var counter int

var mutex sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go increment(&wg)
	go increment(&wg)
	wg.Wait()
	fmt.Println("Final Counter: ", counter)
}
