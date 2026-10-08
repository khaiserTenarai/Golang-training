package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go func() {
		fmt.Println("[Sender] Preparing data...")
		time.Sleep(500 * time.Millisecond)

		fmt.Println("[Sender] Sending message into unbuffered channel...")
		ch <- "Task completed successfully"

		fmt.Println("[Sender] Message received by receiver! Sender proceeding...")
	}()

	fmt.Println("[Receiver] Waiting to receive data...")
	message := <-ch 

	fmt.Println("[Receiver] Received:", message)
}