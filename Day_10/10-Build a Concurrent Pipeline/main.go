package main

import "fmt"

// Stage 1: Generate numbers.
func generate(numbers ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, number := range numbers {
			out <- number
		}
	}()

	return out
}

// Stage 2: Square the numbers.
func square(input <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for number := range input {
			out <- number * number
		}
	}()

	return out
}

// Stage 3: Double the numbers.
func double(input <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for number := range input {
			out <- number * 2
		}
	}()

	return out
}

func main() {
	/*
		Concurrent Pipeline:

		Stage 1        Stage 2        Stage 3

		Generate  ---> Square  ---> Double ---> Result

		Each stage runs in its own goroutine.
	*/

	numbers := generate(1, 2, 3, 4, 5)

	squared := square(numbers)

	doubled := double(squared)

	// Read final results.
	for result := range doubled {
		fmt.Println(result)
	}

	fmt.Println("Pipeline completed")
}