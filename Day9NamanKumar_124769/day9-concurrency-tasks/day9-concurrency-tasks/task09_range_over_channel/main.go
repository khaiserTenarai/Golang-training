package main

import "fmt"

func produceNumbers(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go produceNumbers(ch)

	for value := range ch {
		fmt.Println("Received:", value)
	}
}
