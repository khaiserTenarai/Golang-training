package main

import "fmt"

func producer(ch chan<- int) {
	fmt.Println("[Producer] Generating numbers...")
	for i := 1; i <= 3; i++ {
		ch <- i * 10
	}
	close(ch)
}

func consumer(ch <-chan int) {
	fmt.Println("[Consumer] Reading numbers...")
	for val := range ch {
		fmt.Println("[Consumer] Received:", val)
	}
}

func main() {
	ch := make(chan int)

	go producer(ch)
	consumer(ch)

	fmt.Println("[Main] Program finished.")
}