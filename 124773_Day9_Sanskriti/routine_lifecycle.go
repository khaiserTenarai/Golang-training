package main

import (
	"fmt"
	"time"
)

func employeeWork() {
	fmt.Println("Goroutine started")

	time.Sleep(time.Second)

	fmt.Println("Goroutine finished")
}

func main() {
	fmt.Println("Main started")

	go employeeWork()

	fmt.Println("Goroutine created")

	time.Sleep(2 * time.Second)

	fmt.Println("Main finished")
}