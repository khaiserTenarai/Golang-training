package main

import (
	"fmt"
	"time"
)

func main() {
	// Unbuffered channel synchronization
	ch := make(chan string)

	go func() {
		fmt.Println("Sender: preparing data...")
		time.Sleep(time.Second)
		ch <- "Employee Data Payload" // Blocks until receiver is ready
		fmt.Println("Sender: data sent successfully!")
	}()

	fmt.Println("Receiver: waiting for data...")
	data := <-ch // Blocks until sender writes to the channel
	fmt.Printf("Receiver: received -> %s\n", data)
}