package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("Initialization happens only once")
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	once.Do(initialize) // initialize() runs only once.
	fmt.Println("Worker is running")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(5)

	for i := 1; i <= 5; i++ {
		go worker(&wg)
	}

	wg.Wait()
}
