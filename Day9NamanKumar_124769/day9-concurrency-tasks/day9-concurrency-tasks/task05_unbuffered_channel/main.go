package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "message from goroutine"
	}()

	msg := <-ch
	fmt.Println("Received:", msg)
}
