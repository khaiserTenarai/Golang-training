package main
import (
	"fmt"
	"sync"
)

func main() {
	/*
		Problem:
		Run multiple goroutines and wait until all of them finish.

		Solution:
		sync.WaitGroup keeps track of goroutines.

		Important functions:
		Add()    -> tells how many goroutines are starting
		Done()   -> tells one goroutine has finished
		Wait()   -> waits until all goroutines finish
	*/

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			fmt.Println("Goroutine", id, "completed")
		}(i)
	}

	// Wait until all 3 goroutines complete.
	wg.Wait()

	fmt.Println("All goroutines completed")
}