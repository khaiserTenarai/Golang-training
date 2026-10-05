package main

import "fmt"

// sendOnly can only send data to the channel.
func sendEmployee(ch chan<- string) {
	ch <- "Employee processing completed"
}

// receiveOnly can only receive data from the channel.
func receiveEmployee(ch <-chan string) {
	message := <-ch
	fmt.Println("Received:", message)
}

func main() {
	employeeChannel := make(chan string)

	go sendEmployee(employeeChannel)

	receiveEmployee(employeeChannel)

	fmt.Println("Processing completed")
}
