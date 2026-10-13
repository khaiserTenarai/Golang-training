
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main goroutine started")

	go func() {
		fmt.Println("Anonymous goroutine started")

		for i := 1; i <= 5; i++ {
			fmt.Println("Anonymous goroutine:", i)
			time.Sleep(500 * time.Millisecond)
		}

		fmt.Println("Anonymous goroutine completed")
	}()

	time.Sleep(3 * time.Second)

	fmt.Println("Main goroutine completed")
}