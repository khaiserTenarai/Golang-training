package main

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
)

func initialize() {
	fmt.Println("--> Initialization function executed (Only once!)")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Goroutine %d starting\n", id)

	once.Do(initialize)

	fmt.Printf("Goroutine %d continuing work\n", id)
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 5

	wg.Add(numGoroutines)

	for i := 1; i <= numGoroutines; i++ {
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All goroutines completed.")
}