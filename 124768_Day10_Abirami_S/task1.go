package main

import (
	"fmt"
	"sync"
)

func printMessage(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Goroutine", id, " is running")
}
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go printMessage(i, &wg)
	}
	wg.Wait()
	fmt.Println("All goroutines completed")
}
