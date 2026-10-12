package main

import "fmt" // Imports fmt for output.

func numbers() <-chan int { // Creates the first pipeline stage.
	out := make(chan int) // Creates an output channel.

	go func() { // Starts the producer goroutine.
		defer close(out) // Closes the channel when production finishes.

		for i := 1; i <= 5; i++ { // Generates five numbers.
			out <- i // Sends each number to the next stage.
		}
	}()

	return out // Returns the read-only channel.
}

func multiply(in <-chan int) <-chan int { // Creates the second pipeline stage.
	out := make(chan int) // Creates another output channel.

	go func() { // Starts the multiplication goroutine.
		defer close(out) // Closes output when input ends.

		for n := range in { // Reads values from the previous stage.
			out <- n * 2 // Multiplies each number by 2.
		}
	}()

	return out // Returns the read-only channel.
}

func main() {
	for result := range multiply(numbers()) { // Connects the pipeline stages.
		fmt.Println(result) // Prints each final result.
	}
}
