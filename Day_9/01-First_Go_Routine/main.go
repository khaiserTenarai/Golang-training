package main

import (
	"fmt"
	"time"
)

/*
hello() is a normal function.

When we call:
    go hello()

the function runs as a separate goroutine.
*/
func hello() {
	fmt.Println("Hello from goroutine")
}

func main(){

	/*
		Start hello() as a goroutine.

		"go" tells Go to execute the function
		concurrently instead of waiting for it
		to finish in the normal way.
	*/
	go hello()

	/*
		Sleep for 1 second.

		This gives the hello goroutine time to execute
		before the main goroutine finishes.

		Without Sleep, main() may finish first,
		and the program may terminate before
		hello() gets a chance to print.
	*/
	time.Sleep(time.Second)

	/*
		After waiting for 1 second, the main goroutine
		continues execution and prints this message.
	*/
	fmt.Println("Main completed")
}

/*
Output:
---------
PS C:\Training\HP - GO Language Training\Day_9\124772_Day_9(Go)_Assignment_Reddem_Ganesh_Reddy\01-First_Go_Routine> go run .\main.go
Hello from goroutine
Main completed
*/