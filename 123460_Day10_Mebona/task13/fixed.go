package main

import "fmt"

func main() {
	// Solution 1: Use a buffered channel (capacity 1) so sending doesn't block
	ch := make(chan int, 1)
	ch <- 42
	fmt.Println("Buffered Channel Value:", <-ch)

	// Solution 2: Send from a separate goroutine so the main thread can receive
	chUnbuffered := make(chan int)
	go func() {
		chUnbuffered <- 100
	}()
	fmt.Println("Unbuffered Channel Value:", <-chUnbuffered)
}