package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "Employee 1"
		ch <- "Employee 2"
		ch <- "Employee 3"
		close(ch)
	}()
	for {
		message, ok := <-ch
		if !ok {
			fmt.Println("Channel Closed")
			break
		}
		fmt.Println(message)
	}
}
