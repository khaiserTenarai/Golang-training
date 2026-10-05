package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main function started")

	go func() {
		fmt.Println("Anonymous goroutine started")

		time.Sleep(1 * time.Second)

		fmt.Println("Employee calculation completed")
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("Main function completed")
}
