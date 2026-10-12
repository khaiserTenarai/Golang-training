package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	ch <- 101
	ch <- 102
	ch <- 103
	fmt.Println("Sent 3 employee IDs without any receiver")
	fmt.Println("length:", len(ch), "capacity:", cap(ch))

	for len(ch) > 0 {
		fmt.Println("Received ID:", <-ch)
	}
}
