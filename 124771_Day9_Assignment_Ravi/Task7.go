package main

import "fmt"

func main() {

	ch := make(chan int)
	go func() {

		for i := 1; i <= 8; i++ {
			ch <- i
		}
		close(ch)
	}()

	for value := range ch {
		fmt.Println("Received:", value)
	}

	fmt.Println("Channel completed")

}
