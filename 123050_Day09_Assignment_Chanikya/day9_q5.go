package main

import "fmt"

func main() {
	message := make(chan string)

	go func() {
		fmt.Println("Sending employee message...")
		message <- "Employee processing completed"
	}()

	result := <-message

	fmt.Println("Received:", result)
}
