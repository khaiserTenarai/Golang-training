package main

import "fmt"

// Stage 1: Generate numbers.
func generate(numbers ...int) <-chan int {
	output := make(chan int)

	go func() {
		for _, number := range numbers {
			output <- number
		}
		close(output)
	}()

	return output
}

// Stage 2: Square numbers.
func square(input <-chan int) <-chan int {
	output := make(chan int)

	go func() {
		for number := range input {
			output <- number * number
		}
		close(output)
	}()

	return output
}

func main() {
	// Stage 1.
	numbers := generate(1, 2, 3, 4, 5)

	// Stage 2.
	results := square(numbers)

	// Final stage.
	for result := range results {
		fmt.Println(result)
	}
}
