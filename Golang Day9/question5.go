package main

import (
	"fmt"
	"time"
)

func main() {
	
	messageChan := make(chan string)

	go func() {
		fmt.Println("[Sender] Doing some work.")
		time.Sleep(2 * time.Second) 
		fmt.Println("[Sender] Ready to send data.")
		
		messageChan <- "Secret Password: GoIsFun" 
		
		fmt.Println("[Sender] Hand-off complete! I can continue now.")
	}()

	fmt.Println("[Main] I need data. I will block until the sender provides it.")

	receivedMessage := <-messageChan 

	fmt.Printf("[Main] I received the message: '%s'\n", receivedMessage)
}