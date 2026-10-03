package main

import (
	"fmt"
	"time"
)

func main() {

	messageChan := make(chan string, 3)

	go func() {
		fmt.Println("[Sender] Dropping off Message 1...")
		messageChan <- "Apple"
		
		fmt.Println("[Sender] Dropping off Message 2...")
		messageChan <- "Banana"
		
		fmt.Println("[Sender] Dropping off Message 3...")
		messageChan <- "Cherry"
		
		fmt.Println("[Sender] The buffer is full")
		
	
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("[Main]Time to check the channel.")

	fmt.Printf("[Main] Received: %s\n", <-messageChan)
	fmt.Printf("[Main] Received: %s\n", <-messageChan)
	fmt.Printf("[Main] Received: %s\n", <-messageChan)
}