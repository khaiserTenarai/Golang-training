package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 1; i <= 10; i++ {
		fmt.Println("Producer sending:", i)

		ch <- i

		fmt.Println("Producer sent:", i)
	}

	close(ch)
}

func consumer(ch chan int) {
	for value := range ch {
		fmt.Println("Consumer processing:", value)

		time.Sleep(1 * time.Second)
	}
}

func main() {
	channel := make(chan int, 2)

	go producer(channel)
	go consumer(channel)

	time.Sleep(12 * time.Second)

	fmt.Println("Program completed")
}