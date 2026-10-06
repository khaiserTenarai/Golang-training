package main

import "fmt"

func main() {
	ch := make(chan int) // unbuffered, no receiver
	ch <- 1
	fmt.Println(<-ch)
}
