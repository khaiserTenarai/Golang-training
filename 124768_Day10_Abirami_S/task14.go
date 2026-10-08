package main

import (
	"fmt"
	"sync"
)

var counter int

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		//fix
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

	fmt.Println("Counter:", counter)
}
