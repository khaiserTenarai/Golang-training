package main

import (
	"fmt"
	"time"
)

func producer(employeeChannel chan<- string) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Producer: sending employee", i)

		employeeChannel <- fmt.Sprintf("Employee %d", i)

		fmt.Println("Producer: sent employee", i)
	}

	close(employeeChannel)
}

func consumer(employeeChannel <-chan string) {
	for employee := range employeeChannel {
		fmt.Println("Consumer: processing", employee)

		// Consumer is slower than producer.
		time.Sleep(2 * time.Second)
	}
}

func main() {
	employeeChannel := make(chan string, 2)

	go producer(employeeChannel)

	consumer(employeeChannel)

	fmt.Println("Processing completed")
}
