package main

import (
	"fmt"
	"time"
)

func producer(jobs chan<- int) {
	for i := 1; i <= 10; i++ {
		fmt.Println("Producing job:", i)

		jobs <- i

		fmt.Println("Job sent:", i)
	}

	close(jobs)
}

func consumer(jobs <-chan int) {
	for job := range jobs {
		fmt.Println("Processing job:", job)

		// Consumer is slower
		time.Sleep(1 * time.Second)

		fmt.Println("Completed job:", job)
	}
}

func main() {
	// Buffer can hold only 2 jobs
	jobs := make(chan int, 2)

	go producer(jobs)
	consumer(jobs)

	fmt.Println("All jobs completed")
}
