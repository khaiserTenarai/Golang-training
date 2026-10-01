package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from Goroutine!")
}

func main() {
	go sayHello()
	time.Sleep(100 * time.Millisecond) // Waiting for goroutine to execute
	fmt.Println("Hello from Main!")
}