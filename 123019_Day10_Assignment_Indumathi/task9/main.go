package main

import (
	"fmt"
	"sync"
)

func produce(id int, nums []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1 := produce(1, []int{1, 2, 3})
	ch2 := produce(2, []int{10, 20, 30})

	merged := fanIn(ch1, ch2)
	for val := range merged {
		fmt.Println("Merged output:", val)
	}
}