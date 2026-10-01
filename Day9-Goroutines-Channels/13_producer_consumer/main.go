package main

import (
	"fmt"
	"sync"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("Produced: %d\n", i)
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range ch {
		fmt.Printf("Consumed: %d\n", item)
	}
}

func main() {
	ch := make(chan int, 3)
	var wg sync.WaitGroup

	wg.Add(2)
	go producer(ch, &wg)
	go consumer(ch, &wg)

	wg.Wait()
}