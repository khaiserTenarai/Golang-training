package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {

	for i := 1; i <= 5; i++ {
		fmt.Println("Produced:", i)
		ch <- i
	}

	close(ch)
}

func consumer(ch <-chan int) {

	for value := range ch {
		fmt.Println("Consumed:", value)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {

	ch := make(chan int, 2)

	go producer(ch)

	consumer(ch)

	fmt.Println("Finished")
}