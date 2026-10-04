package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Printf("Anonymous goroutine %d running\n", n)
		}(i)
	}

	wg.Wait()
	fmt.Println("All anonymous goroutines finished")
}
