package main

import "fmt"

func sendNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}

	close(ch)
}

func main() {
	numbers := make(chan int)

	go sendNumbers(numbers)

	for value := range numbers {
		fmt.Println("Received:", value)
	}

	fmt.Println("All values received")
}