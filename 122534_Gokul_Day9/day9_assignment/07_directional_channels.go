package main

import "fmt"

func sendIDs(out chan<- int) {
	for id := 101; id <= 105; id++ {
		out <- id
	}
	close(out)
}

func printIDs(in <-chan int, done chan<- bool) {
	for id := range in {
		fmt.Println("Received employee ID:", id)
	}
	done <- true
}

func main() {
	ids := make(chan int)
	done := make(chan bool)

	go sendIDs(ids)
	go printIDs(ids, done)

	<-done
}
