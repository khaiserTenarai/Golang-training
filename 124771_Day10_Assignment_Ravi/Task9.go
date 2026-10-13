package main

import (
	"fmt"
	"sync"
)

func generate(start int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := start; i < start+3; i++ {
		ch <- i
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	output := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)
	go generate(1, ch1, &wg)
	go generate(4, ch2, &wg)
	go func() {
		wg.Wait()
		close(ch1)
		close(ch2)
	}()

	// Fan-in: Combine multiple channels into one
	go func() {
		var mergeWG sync.WaitGroup
		merge := func(ch <-chan int) {
			defer mergeWG.Done()
			for n := range ch {
				output <- n
			}
		}

		mergeWG.Add(2)
		go merge(ch1)
		go merge(ch2)
		mergeWG.Wait()
		close(output)
	}()

	for n := range output {
		fmt.Println(n)
	}
}
