package main

import "fmt"

func main() {

	ch := make(chan string)

	go func() {
		ch <- "Employee added"
	}()

	message := <-ch

	fmt.Println(message)
}