package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Producer produced:", i)

		ch <- i

		time.Sleep(500 * time.Millisecond)
	}

	close(ch)
}

func consumer(ch chan int) {
	for value := range ch {
		fmt.Println("Consumer consumed:", value)

		time.Sleep(1 * time.Second)
	}
}

func main() {
	channel := make(chan int, 2)

	go producer(channel)
	go consumer(channel)

	time.Sleep(7 * time.Second)

	fmt.Println("Producer-Consumer completed")
}