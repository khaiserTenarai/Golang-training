package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from a goroutine!")
}

func main() {
	// Start the function concurrently using the 'go' keyword
	go sayHello()

	// Wait briefly to allow the goroutine to finish executing before main exits
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Main function finished.")
}