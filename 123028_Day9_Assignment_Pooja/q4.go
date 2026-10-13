package main

import (
	"fmt"
)

func main() {
	message := make(chan string)

	go func() {
		message <- "Hello from Goroutine"
	}()

	result := <-message

	fmt.Println(result)
}