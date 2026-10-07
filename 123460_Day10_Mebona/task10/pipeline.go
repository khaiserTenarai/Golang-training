package main

import (
	"fmt"
)

// Stage 1: Generate numbers
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// Stage 2: Square numbers
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	// Pipeline wiring: generator -> square -> main
	genCh := generator(2, 3, 4, 5)
	sqCh := square(genCh)

	for result := range sqCh {
		fmt.Println("Pipeline output:", result)
	}
}