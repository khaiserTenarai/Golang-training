package main

import (
	"fmt"
	"sync"
)

func generate(start int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := start; i < start+5; i++ {
			out <- i
		}
	}()

	return out
}

func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)

		go func(c <-chan int) {
			defer wg.Done()

			for value := range c {
				out <- value
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1 := generate(1)
	ch2 := generate(100)

	for value := range fanIn(ch1, ch2) {
		fmt.Println(value)
	}
}
