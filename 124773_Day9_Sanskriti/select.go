package main

import "fmt"

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "Employee channel"
	}()

	go func() {
		ch2 <- "Salary channel"
	}()

	select {
	case message := <-ch1:
		fmt.Println(message)

	case message := <-ch2:
		fmt.Println(message)
	}
}