package main

import (
	"fmt"
	"sync"
)

func produce(ch chan<- int, count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= count; i++ {
		ch <- i
	}
}

func consume(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for value := range ch {
		fmt.Println("Consumed:", value)
	}
}

func main() {
	ch := make(chan int, 5)

	var producerWg sync.WaitGroup
	var consumerWg sync.WaitGroup

	producerWg.Add(1)
	go produce(ch, 10, &producerWg)

	consumerWg.Add(1)
	go consume(ch, &consumerWg)

	producerWg.Wait()
	close(ch)

	consumerWg.Wait()
}
