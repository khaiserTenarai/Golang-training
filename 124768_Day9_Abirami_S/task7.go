package main

import "fmt"

func sendEmployee(ch chan<- string) {
	ch <- "Employees data sent"
}
func receiveEmployee(ch <-chan string) {
	message := <-ch
	fmt.Println(message)
}
func main() {
	ch := make(chan string)
	go sendEmployee(ch)
	receiveEmployee(ch)
}
