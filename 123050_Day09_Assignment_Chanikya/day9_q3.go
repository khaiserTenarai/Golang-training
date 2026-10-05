package main

import (
	"fmt"
	"time"
)

func employeeTask() {
	fmt.Println("Goroutine started")

	time.Sleep(1 * time.Second)

	fmt.Println("Goroutine is running")

	time.Sleep(1 * time.Second)

	fmt.Println("Goroutine completed")
}

func main() {
	fmt.Println("Main: Creating goroutine")

	go employeeTask()

	fmt.Println("Main: Goroutine started")

	time.Sleep(3 * time.Second)

	fmt.Println("Main: Program completed")
}
