package main

import (
	"fmt"
	"sync"
)

// Fan-In merges multiple input channels into a single output channel
func fanIn(channels ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup

	for _, c := range channels {
		wg.Add(1)
		go func(ch <-chan string) {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func generateSource(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 1; i <= 3; i++ {
			ch <- fmt.Sprintf("%s - %d", msg, i)
		}
	}()
	return ch
}

func main() {
	ch1 := generateSource("Source A")
	ch2 := generateSource("Source B")

	merged := fanIn(ch1, ch2)

	for val := range merged {
		fmt.Println("Received:", val)
	}
}