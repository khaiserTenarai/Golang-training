package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		Problem:
		We want some code to execute only ONE time,
		even when many goroutines call it.

		Solution:
		sync.Once

		once.Do(function)

		The function passed to Do() runs only once.
	*/

	var once sync.Once
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			once.Do(func() {
				fmt.Println("Initialization executed by goroutine", id)
			})

			fmt.Println("Goroutine", id, "finished")
		}(i)
	}

	wg.Wait()
}