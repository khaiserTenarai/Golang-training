package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("Producing: ", i)
			ch <- i
			fmt.Println("Produced: ", i)
		}
		close(ch)
	}()
	for value := range ch {
		fmt.Println("Consuming: ", value)
		time.Sleep(2 * time.Second)
	}
	fmt.Println("All values processed")
}
