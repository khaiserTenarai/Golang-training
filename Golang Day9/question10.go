package main

import (
	"fmt"
	"time"
)

func main() {
	fastChan := make(chan string)
	slowChan := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		fastChan <- "Response from Fast Server"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		slowChan <- "Response from Slow Server"
	}()

	fmt.Println("Waiting for responses...")

	for i := 0; i < 2; i++ {
		select {
		case msg := <-fastChan:
			fmt.Printf("Received: %s\n", msg)
		case msg := <-slowChan:
			fmt.Printf("Received: %s\n", msg)
		}
	}

	fmt.Println("All responses received. Exiting.")
}