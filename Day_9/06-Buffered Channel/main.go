package main
import "fmt"

func main() {

	/*
		Buffered channel can store 3 values.
	*/
	ch := make(chan int, 3)

	/*
		This creates a goroutine.
		The goroutine sends values into the channel.
	*/
	go func() {

		ch <- 10
		ch <- 20
		ch <- 30

	}()

	/*
		Main goroutine receives the values.
	*/
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("Main completed")
}

/*

PS C:\Training\HP - GO Language Training\Day_9\124772_Day_9(Go)_Assignment_Reddem_Ganesh_Reddy\06-Buffered Channel> go run .\main.go
10
20
30
Main completed
*/