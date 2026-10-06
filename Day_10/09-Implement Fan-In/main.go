package main

import (
	"fmt"
	"sync"
)
// Fan-in means combining multiple channels into one channel.
func generate(numbers []int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, number := range numbers {
			out <- number
		}
	}()

	return out
}

func fanIn(ch1 <-chan int, ch2 <-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)

	// Read from first channel.
	go func() {
		defer wg.Done()

		for value := range ch1 {
			out <- value
		}
	}()

	// Read from second channel.
	go func() {
		defer wg.Done()

		for value := range ch2 {
			out <- value
		}
	}()

	// Close output after both channels finish.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	
	ch1 := generate([]int{1, 2, 3})
	ch2 := generate([]int{4, 5, 6})

	result := fanIn(ch1, ch2)

	for value := range result {
		fmt.Println(value)
	}

	fmt.Println("Fan-in completed")
}
// Fan-in = Multiple sources → One channel
