package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			fmt.Println("Worker", id, "finished")
		}(i)
	}

	wg.Wait()

	fmt.Println("All workers finished")
}
