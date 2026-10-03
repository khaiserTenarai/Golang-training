package main

import (
	"fmt"
	"time"
)

func printMessage() {
	fmt.Println("Hello from Goroutine")
}
func main() {
	go printMessage()
	time.Sleep(1 * time.Second)
	fmt.Println("Main function completed")
}
