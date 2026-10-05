package main

import (
	"fmt"
	"sync"
)

func merge(
	ch1 <-chan string,
	ch2 <-chan string,
) <-chan string {

	output := make(chan string)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		for value := range ch1 {
			output <- value
		}
	}()

	go func() {
		defer wg.Done()

		for value := range ch2 {
			output <- value
		}
	}()

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "Employee 1"
		ch1 <- "Employee 2"
		close(ch1)
	}()

	go func() {
		ch2 <- "Employee 3"
		ch2 <- "Employee 4"
		close(ch2)
	}()

	// Combine two channels into one
	output := merge(ch1, ch2)

	for value := range output {
		fmt.Println(value)
	}
}