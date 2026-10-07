package main

import (
	"context"
	"fmt"
	"time"
)

func fetchRemoteData() chan string {
	ch := make(chan string)
	go func() {
		time.Sleep(2 * time.Second) // Simulating slow response
		ch <- "data payload"
	}()
	return ch
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	dataCh := fetchRemoteData()

	select {
	case res := <-dataCh:
		fmt.Println("Received:", res)
	case <-ctx.Done():
		fmt.Println("Operation timed out:", ctx.Err())
	}
}