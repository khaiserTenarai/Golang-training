package main

import "fmt"

func generate() <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := 1; i <= 5; i++ {
			out <- i
		}
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for value := range in {
			out <- value * value
		}
	}()

	return out
}

func double(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for value := range in {
			out <- value * 2
		}
	}()

	return out
}

func main() {
	numbers := generate()
	squares := square(numbers)
	results := double(squares)

	for result := range results {
		fmt.Println(result)
	}
}
