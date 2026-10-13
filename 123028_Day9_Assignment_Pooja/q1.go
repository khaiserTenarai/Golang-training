package main

import (
	"fmt"
	"time"
)

func printMessage() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Hello from Goroutine", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printMessage()

	for i := 1; i <= 5; i++ {
		fmt.Println("Hello from Main", i)
		time.Sleep(500 * time.Millisecond)
	}
}