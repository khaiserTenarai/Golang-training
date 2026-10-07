package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			fmt.Printf("Worker %d started\n", id)
			time.Sleep(time.Second)
			fmt.Printf("Worker %d finished\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("All workers finished")
}
