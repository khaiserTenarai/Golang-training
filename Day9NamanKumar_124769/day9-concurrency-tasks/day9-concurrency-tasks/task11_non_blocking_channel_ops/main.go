package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	select {
	case value := <-ch:
		fmt.Println("Received:", value)
	default:
		fmt.Println("No value available, moving on")
	}

	select {
	case ch <- 10:
		fmt.Println("Sent value")
	default:
		fmt.Println("Channel full, skipping send")
	}

	select {
	case value := <-ch:
		fmt.Println("Received:", value)
	default:
		fmt.Println("No value available, moving on")
	}
}
