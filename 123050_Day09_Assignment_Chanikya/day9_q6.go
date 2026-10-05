package main

import "fmt"

func main() {
	messages := make(chan string, 3)

	messages <- "Employee 1 completed"
	messages <- "Employee 2 completed"
	messages <- "Employee 3 completed"

	fmt.Println("Buffered channel contains 3 messages")

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
