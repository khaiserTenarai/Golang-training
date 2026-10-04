package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)

	ch <- 100
	ch <- 200
	ch <- 300
	close(ch) // Close channel when no more sends are coming

	// Reading values and checking if channel is open
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel is closed. Terminating read loop.")
			break
		}
		fmt.Printf("Read value: %d\n", val)
	}
}