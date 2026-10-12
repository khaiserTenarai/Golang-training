package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	select {
	case v := <-ch:
		fmt.Println("Received:", v)
	default:
		fmt.Println("Nothing to receive, moving on")
	}

	select {
	case ch <- 5:
		fmt.Println("Sent 5")
	default:
		fmt.Println("Channel full, could not send 5")
	}

	select {
	case ch <- 6:
		fmt.Println("Sent 6")
	default:
		fmt.Println("Channel full, could not send 6")
	}

	select {
	case v := <-ch:
		fmt.Println("Received:", v)
	default:
		fmt.Println("Nothing to receive, moving on")
	}
}
