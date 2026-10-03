package main

import "fmt"

/*
	DIRECTIONAL CHANNELS

	chan int
		→ Can send and receive

	chan<- int
		→ Send only

	<-chan int
		→ Receive only
*/

// This function can ONLY send to the channel.
func sendData(ch chan<- int) {

	fmt.Println("Sending data...")

	ch <- 100
}

// This function can ONLY receive from the channel.
func receiveData(ch <-chan int) {

	value := <-ch

	fmt.Println("Received:", value)
}

func main() {

	// Normal channel: can send and receive.
	ch := make(chan int)

	// Create a goroutine for sending.
	go sendData(ch)

	// Receive data.
	receiveData(ch)
}

/*
Output:
----------
PS C:\Training\HP - GO Language Training\Day_9\124772_Day_9(Go)_Assignment_Reddem_Ganesh_Reddy\07-Directional Channels> go run .\main.go
Sending data...
Received: 100
*/