package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "task completed"
	}()

	select {
	case res := <-ch:
		fmt.Println("Result:", res)
	case <-time.After(2 * time.Second):
		fmt.Println("Error: task timed out")
	}
}