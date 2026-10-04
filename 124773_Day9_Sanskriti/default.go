package main

import "fmt"

func main() {

	ch := make(chan string)

	select {
	case message := <-ch:
		fmt.Println(message)

	default:
		fmt.Println("No message available")
	}
}