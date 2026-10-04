package main

import "fmt"

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		channel1 <- "Message from Channel 1"
	}()

	go func() {
		channel2 <- "Message from Channel 2"
	}()

	select {
	case message1 := <-channel1:
		fmt.Println(message1)

	case message2 := <-channel2:
		fmt.Println(message2)
	}
}