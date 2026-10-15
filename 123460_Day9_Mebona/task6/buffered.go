package main

import (
	"fmt"
)

func main() {
	// Buffered channel with a capacity of 3
	ch := make(chan string, 3)

	// Does not block because capacity is not exceeded
	ch <- "Task 1"
	ch <- "Task 2"
	ch <- "Task 3"

	fmt.Printf("Channel buffer length: %d, capacity: %d\n", len(ch), cap(ch))

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}