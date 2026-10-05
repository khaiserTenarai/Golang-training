package main

import (
	"fmt"
	"time"
)

func main() {
	employee1 := make(chan string)
	employee2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		employee1 <- "Employee 1 completed"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		employee2 <- "Employee 2 completed"
	}()

	for i := 0; i < 2; i++ {
		select {
		case message := <-employee1:
			fmt.Println("Received:", message)

		case message := <-employee2:
			fmt.Println("Received:", message)
		}
	}

	fmt.Println("All employees processed")
}
