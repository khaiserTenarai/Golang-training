package main

import (
	"fmt"
	"time"
)

func producer(jobs chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Produced:", i)
		jobs <- i
		time.Sleep(500 * time.Millisecond)
	}

	close(jobs)
}

func consumer(jobs <-chan int) {
	for job := range jobs {
		fmt.Println("Consumed:", job)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	jobs := make(chan int, 2)

	go producer(jobs)

	consumer(jobs)

	fmt.Println("Done")
}
