package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker4(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped:", ctx.Err())
			return
		default:
			fmt.Println("Worker processing:", i)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)
	go worker4(ctx, &wg)
	wg.Wait()
	fmt.Println("Program completed")
}
