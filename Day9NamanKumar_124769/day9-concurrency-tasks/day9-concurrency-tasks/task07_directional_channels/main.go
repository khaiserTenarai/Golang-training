package main

import "fmt"

func send(ch chan<- int, value int) {
	ch <- value
}

func receive(ch <-chan int) int {
	return <-ch
}

func main() {
	ch := make(chan int, 1)
	send(ch, 42)
	value := receive(ch)
	fmt.Println("Received:", value)
}
