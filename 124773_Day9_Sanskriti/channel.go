package main

import "fmt"

func sendEmployee(ch chan<- string) {
	ch <- "Sanskriti"
}

func receiveEmployee(ch <-chan string) {
	fmt.Println(<-ch)
}

func main() {

	ch := make(chan string)

	go sendEmployee(ch)

	receiveEmployee(ch)
}