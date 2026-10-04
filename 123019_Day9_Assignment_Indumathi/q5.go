package main

import (
	"fmt"
)

func main() {
	message := make(chan string, 3)

	message <- "Hello"
	message <- "Welcome"
	message <- "Go Programming"

	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}