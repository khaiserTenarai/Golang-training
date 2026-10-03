package main
import (
	"fmt"
	"time"
)

/*
	PRODUCER-CONSUMER

	Producer
	    |
	    | sends data
	    ↓
	 Channel
	    |
	    | receives data
	    ↓
	Consumer

	Producer → creates values
	Consumer → processes values
	Channel  → connects them
*/

func producer(ch chan int) {

	// Produce 5 values.
	for i := 1; i <= 5; i++ {

		fmt.Println("Producer: Sending", i)

		ch <- i

		// Small delay between producing values.
		time.Sleep(500 * time.Millisecond)
	}

	// Producer has no more data.
	close(ch)
}

func consumer(ch chan int) {

	/*
		range receives values until the channel
		is closed by the producer.
	*/
	for value := range ch {

		fmt.Println("Consumer: Received", value)

		// Consumer takes some time to process data.
		time.Sleep(1 * time.Second)
	}
}

func main() {

	// Create a buffered channel.
	ch := make(chan int, 2)

	// Start producer goroutine.
	go producer(ch)

	// Start consumer goroutine.
	go consumer(ch)

	/*
		Keep main alive so both goroutines
		get time to complete.
	*/
	time.Sleep(6 * time.Second)

	fmt.Println("Main completed")
}

/*
Output:
--------
PS C:\Training\HP - GO Language Training\Day_9\124772_Day_9(Go)_Assignment_Reddem_Ganesh_Reddy\12-Producer–Consumer> go run .\main.go
Producer: Sending 1
Consumer: Received 1
Producer: Sending 2
Consumer: Received 2
Producer: Sending 3
Producer: Sending 4
Consumer: Received 3
Producer: Sending 5
Consumer: Received 4
Consumer: Received 5
Main completed
*/