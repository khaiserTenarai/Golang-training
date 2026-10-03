package main
import "fmt"

func main() {

	/*
		Unbuffered Channel

		make(chan int) creates an unbuffered channel.

		Capacity = 0

		Sender  --->  Channel  --->  Receiver

		The sender waits until the receiver receives
		the value.
	*/

	ch := make(chan int)

	// Send value from a goroutine
	go func() {
		fmt.Println("Sending value...")
		ch <- 100
		fmt.Println("Value sent")
	}()

	// Receive value
	value := <-ch

	fmt.Println("Received:", value)
}

/*
Output :
---------
PS C:\Training\HP - GO Language Training\Day_9\124772_Day_9(Go)_Assignment_Reddem_Ganesh_Reddy\05-Unbuffered Channel> go run .\main.go
Sending value...
Value sent
Received: 100
*/