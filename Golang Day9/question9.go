package main

import "fmt"

func main() {
	
	words := make(chan string, 4)

	words <- "Go"
	words <- "channels"
	words <- "are"
	words <- "awesome!"

	close(words)

	fmt.Println("Starting to read from the channel...")

	for word := range words {
		fmt.Printf("Received: %s\n", word)
	}
	fmt.Println("Loop finished automatically because the channel was closed.")
}