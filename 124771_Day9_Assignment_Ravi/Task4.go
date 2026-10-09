package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {

		ch <- 100

	}()

	value := <-ch

	fmt.Println("Received:", value)
}
