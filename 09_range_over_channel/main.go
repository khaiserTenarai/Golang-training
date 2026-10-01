package main

import "fmt"

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch) // Must close so range loop terminates
}

func main() {
	ch := make(chan int)
	go producer(ch)

	for num := range ch {
		fmt.Println("Received:", num)
	}
}