package main
import "fmt"

func main() {

	/*
		CHANNEL CLOSING

		close(ch) closes the channel.

		After closing:
		- We cannot send new values.
		- We can still receive remaining values.
	*/

	ch := make(chan int)

	// Goroutine sends values.
	go func() {

		ch <- 10
		ch <- 20
		ch <- 30

		// No more values will be sent.
		close(ch)
	}()

	/*
		"range" receives values until the channel is closed.
	*/
	for value := range ch {
		fmt.Println("Received:", value)
	}

	fmt.Println("Channel closed")
}

/*
Output:
-------------
PS C:\Training\HP - GO Language Training\Day_9\124772_Day_9(Go)_Assignment_Reddem_Ganesh_Reddy\08-Channel Closing> go run main.go
Received: 10
Received: 20
Received: 30
Channel closed

*/