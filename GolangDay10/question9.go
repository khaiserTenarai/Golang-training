package main

import (
	"fmt"
	"sync"
)

func producer(start int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			out <- start + i
		}
		close(out)
	}()
	return out
}

func fanIn(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1 := producer(10)
	ch2 := producer(20)
	ch3 := producer(30)

	merged := fanIn(ch1, ch2, ch3)

	for val := range merged {
		fmt.Println(val)
	}
}