package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		SELECT

		select is similar to switch,
		but it is used with channels.

		It waits for multiple channel operations
		and executes ONE case that is ready first.

		IMPORTANT:
		- select executes ONLY ONE case.
		- After one case executes, the select statement is finished.
		- It does NOT automatically execute the other case.
	*/

	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from Channel 1"
	}()

	// Goroutine 2
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from Channel 2"
	}()

	/*
		After 1 second:
		ch1 becomes ready.

		So this case executes:
			case msg := <-ch1

		The select then ENDS.

		Even though ch2 sends a message after 2 seconds,
		the ch2 case will NOT execute because select has
		already finished.
	*/
	// for i := 0; i < 2; i++{
	select {

	case msg := <-ch1:
		fmt.Println(msg)

	case msg := <-ch2:
		fmt.Println(msg)
	// }

	/*
		default:
			fmt.Println("No data available")

		If no channel is ready, default executes immediately.
	*/
	}

	/*
		IMPORTANT:

		This Sleep does NOT execute the second select case.

		It only keeps the main goroutine alive for 3 seconds.

		ch2 will send its message after 2 seconds,
		but there is NO receiver waiting for it anymore.

		Therefore "Message from Channel 2" is NOT printed.
	*/
	time.Sleep(3 * time.Second)

	fmt.Println("Main completed")
}
/*
Output:
------
PS C:\Training\HP - GO Language Training\Day_9\124772_Day_9(Go)_Assignment_Reddem_Ganesh_Reddy\10-Select for Channels> go run .\main.go
Message from Channel 1
Main completed
*/