package main

import (
	"fmt"
	"sync"
)

func main() {

	var once sync.Once
	var wg sync.WaitGroup

	wg.Add(5)

	for i := 1; i <= 5; i++ {

		go func(id int) {
			defer wg.Done()

			once.Do(func() {
				fmt.Println("Initialization executed by worker", id)
			})

		}(i)
	}

	wg.Wait()

	fmt.Println("Finished")
}