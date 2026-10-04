package main

import "fmt"

func sendMessage(ch chan<- string) {
	ch <- "Hello from send-only channel"
}

func receiveMessage(ch <-chan string) {
	message := <-ch
	fmt.Println(message)
}

func main() {
	message := make(chan string)

	go sendMessage(message)

	receiveMessage(message)
}