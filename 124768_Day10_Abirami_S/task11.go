package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker3(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Worker cancelled")
			return
		default:
			fmt.Println("Worker processing:", i)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go worker3(ctx, &wg)
	time.Sleep(2 * time.Second)
	fmt.Println("Cancelling worker...")
	cancel()
	wg.Wait()
	fmt.Println("Program completed")
}
