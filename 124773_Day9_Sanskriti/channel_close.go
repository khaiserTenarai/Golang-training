package main

import "fmt"

func main() {

	ch := make(chan string)

	go func() {
		ch <- "Employee 1"
		ch <- "Employee 2"

		close(ch)
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}