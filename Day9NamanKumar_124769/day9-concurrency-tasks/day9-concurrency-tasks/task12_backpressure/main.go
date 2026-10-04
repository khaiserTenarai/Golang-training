package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		fmt.Println("Producing:", i)
		ch <- i
		fmt.Println("Produced:", i)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for value := range ch {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Consumed:", value)
	}
}

func main() {
	ch := make(chan int, 2)
	done := make(chan struct{})

	go producer(ch)
	go func() {
		consumer(ch)
		close(done)
	}()

	<-done
}
