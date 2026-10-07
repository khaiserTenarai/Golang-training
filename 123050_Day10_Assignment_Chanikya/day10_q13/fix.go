package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu1 sync.Mutex
	var mu2 sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2)

	work := func(id int) {
		defer wg.Done()

		mu1.Lock()
		defer mu1.Unlock()

		mu2.Lock()
		defer mu2.Unlock()

		fmt.Println("Goroutine", id)
	}

	go work(1)
	go work(2)

	wg.Wait()
}
