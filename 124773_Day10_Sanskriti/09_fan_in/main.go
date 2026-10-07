package main

import (
	"fmt"
	"sync"
)

func employee1() <-chan string {
	ch := make(chan string)

	go func() {
		ch <- "Employee 1 completed"
		close(ch)
	}()

	return ch
}

func employee2() <-chan string {
	ch := make(chan string)

	go func() {
		ch <- "Employee 2 completed"
		close(ch)
	}()

	return ch
}

// fanIn combines multiple input channels into one output channel.
func fanIn(ch1, ch2 <-chan string) <-chan string {
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
	ch1 := employee1()
	ch2 := employee2()

	results := fanIn(ch1, ch2)

	for result := range results {
		fmt.Println(result)
	}
}
