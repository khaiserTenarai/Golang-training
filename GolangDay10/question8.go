package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Printf("Consumer %d received value %d\n", id, val)
		time.Sleep(time.Millisecond * 50)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	go producer(ch)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go consumer(i, ch, &wg)
	}

	wg.Wait()
}