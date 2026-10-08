package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 1)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	default:
		fmt.Println("No message received (channel was empty)")
	}

	messages <- "Hello Go"

	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	default:
		fmt.Println("No message received")
	}

	msg := "High Priority Task"
	select {
	case messages <- msg:
		fmt.Println("Sent message:", msg)
	default:
		fmt.Println("Could not send message (channel is full/unbuffered without receiver)")
	}

	messages <- "Task 1"

	select {
	case messages <- "Task 2":
		fmt.Println("Sent Task 2")
	default:
		fmt.Println("Buffer full! Skipped sending Task 2 to avoid blocking")
	}

	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	case sig := <-signals:
		fmt.Println("Received signal:", sig)
	default:
		fmt.Println("No activity across any channels")
	}
}