package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Worker %d done\n", id)

}

func main() {
	var wg sync.WaitGroup

	numWorkers := 5

	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg)
	}

	wg.Wait()

	fmt.Println("All workers completed successfully.")
}