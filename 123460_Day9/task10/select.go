package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Payroll Database Response"
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "HR Portal Response"
	}()

	// Select handles whichever channel responds first
	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	}
}