package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println("Sender: about to send")
		ch <- "Employee Asha has joined"
		fmt.Println("Sender: message was received")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Main: ready to receive")
	msg := <-ch
	fmt.Println("Main got:", msg)

	time.Sleep(100 * time.Millisecond)
}
