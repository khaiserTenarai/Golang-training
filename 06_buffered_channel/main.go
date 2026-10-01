package main

import "fmt"

func main() {
	ch := make(chan string, 2) // Capacity of 2

	// Sends without blocking because buffer is not full
	ch <- "Message 1"
	ch <- "Message 2"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}