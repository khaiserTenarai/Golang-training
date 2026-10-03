package main

import (
	"fmt"
	"time"
)

func main() {

	paymentChan := make(chan string)
	notificationChan := make(chan string)

	go func() {
		time.Sleep(300 * time.Millisecond)
		paymentChan <- "Payment #1001 Processed"
	}()

	go func() {
		time.Sleep(600 * time.Millisecond)
		notificationChan <- "Notification Sent to User"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-paymentChan:
			fmt.Println("[Select Received]:", msg1)
		case msg2 := <-notificationChan:
			fmt.Println("[Select Received]:", msg2)
		}
	}

	slowChan := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		slowChan <- "Slow Response"
	}()

	fmt.Println("\nWaiting for slow operation with 1-second timeout...")
	select {
	case res := <-slowChan:
		fmt.Println("[Received]:", res)
	case <-time.After(1 * time.Second):
		fmt.Println("[Timeout]: Operation took too long and timed out!")
	}
}