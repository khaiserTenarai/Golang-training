package main

import "fmt"

func main() {
	channel := make(chan string)

	select {
	case message := <-channel:
		fmt.Println("Received:", message)

	default:
		fmt.Println("No message available")
	}

	fmt.Println("Program completed")
}