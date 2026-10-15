package main

import (
	"fmt"
)

// Send-only channel parameter
func produceData(ch chan<- string) {
	ch <- "Performance Metric A"
	ch <- "Performance Metric B"
}

// Receive-only channel parameter
func consumeData(ch <-chan string) {
	fmt.Println("Consumed:", <-ch)
	fmt.Println("Consumed:", <-ch)
}

func main() {
	ch := make(chan string, 2)

	produceData(ch)
	consumeData(ch)
}