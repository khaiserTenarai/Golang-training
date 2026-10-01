package main

import "fmt"

// Send-only channel
func sendData(ch chan<- string) {
	ch <- "Hello from send-only channel"
}

// Receive-only channel
func receiveData(ch <-chan string) {
	fmt.Println("Received:", <-ch)
}

func main() {
	ch := make(chan string, 1)
	sendData(ch)
	receiveData(ch)
}