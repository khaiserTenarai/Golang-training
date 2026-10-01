package main

import "fmt"

func sendData(ch chan string) {
	ch <- "Data sent through unbuffered channel"
}

func main() {
	ch := make(chan string) // Unbuffered

	go sendData(ch)

	msg := <-ch // Blocks until data is received
	fmt.Println("Received:", msg)
}