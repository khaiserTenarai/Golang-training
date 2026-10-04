package main

import "fmt"

func main() {

	ch := make(chan string)

	go func() {
		ch <- "Sanskriti"
		ch <- "Rahul"
		ch <- "Priya"

		close(ch)
	}()

	for employee := range ch {
		fmt.Println(employee)
	}
}