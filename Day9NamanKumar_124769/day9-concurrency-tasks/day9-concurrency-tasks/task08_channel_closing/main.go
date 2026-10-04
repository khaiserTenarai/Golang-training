package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("Channel closed, no more values")
			break
		}
		fmt.Println("Received:", value)
	}
}
