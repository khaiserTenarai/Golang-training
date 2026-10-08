package main

import (
	"fmt"
	"sync"
)

func worker2(id int, output chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	result := fmt.Sprintf("Result from Worker %d ", id)
	output <- result
}
func fanIn(ch1 <-chan string, ch2 <-chan string, ch3 <-chan string) <-chan string {
	output := make(chan string)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		for value := range ch1 {
			output <- value
		}
	}()
	go func() {
		defer wg.Done()
		for value := range ch2 {
			output <- value
		}
	}()
	go func() {
		defer wg.Done()
		for value := range ch3 {
			output <- value
		}
	}()
	go func() {
		wg.Wait()
		close(output)
	}()
	return output
}
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	var wg sync.WaitGroup
	wg.Add(3)
	go worker2(1, ch1, &wg)
	go worker2(2, ch2, &wg)
	go worker2(3, ch3, &wg)

	go func() {
		wg.Wait()
		close(ch1)
		close(ch2)
		close(ch3)
	}()
	output := fanIn(ch1, ch2, ch3)
	for result := range output {
		fmt.Println(result)
	}
	fmt.Println("All results received")
}
