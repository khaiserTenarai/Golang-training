package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20
	close(ch)

	val1, ok1 := <-ch
	fmt.Printf("Val: %d, Open: %v\n", val1, ok1)

	val2, ok2 := <-ch
	fmt.Printf("Val: %d, Open: %v\n", val2, ok2)

	val3, ok3 := <-ch // Reading from closed channel returns zero value and false
	fmt.Printf("Val: %d, Open: %v\n", val3, ok3)
}