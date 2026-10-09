package main

import "fmt"

func producer(ch chan<- int) {
	ch <- 1000
	close(ch)
}

func consumer(ch <-chan int) {
	value := <-ch
	fmt.Println("Received:", value)
}

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}
