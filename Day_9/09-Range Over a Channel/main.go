package main

import "fmt"

func main() {

	/*
		RANGE OVER CHANNEL

		for value := range ch

		receives values from the channel one by one.

		The loop automatically stops when the channel
		is closed.
	*/

	ch := make(chan int)

	// Goroutine sends values.
	go func() {

		ch <- 10
		ch <- 20
		ch <- 30

		// Tell the receiver that no more values
		// will be sent.
		close(ch)

	}()

	// Receive values using range.
	for value := range ch {

		fmt.Println("Received:", value)
	}

	fmt.Println("Channel closed")
}

/*
Output:
------------
PS C:\Training\HP - GO Language Training\Day_9\124772_Day_9(Go)_Assignment_Reddem_Ganesh_Reddy\09-Range Over a Channel> go run .\main.go
Received: 10
Received: 20
Received: 30
Channel closed
*/