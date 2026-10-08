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
	for employee := range ch {
		fmt.Println(employee)
	}
	fmt.Println("All employees processed")
}
