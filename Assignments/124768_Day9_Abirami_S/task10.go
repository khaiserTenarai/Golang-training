package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from Channel 1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from Channel 2"
	}()
	for i := 0; i < 2; i++ {
		select {
		case message := <-ch1:
			fmt.Println(message)
		case message := <-ch2:
			fmt.Println(message)
		}
	}
	fmt.Println("All messages received")
}
