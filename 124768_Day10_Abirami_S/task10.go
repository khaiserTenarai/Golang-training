package main

import (
	"fmt"
	"sync"
)

func generate(numbers chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(numbers)

	for i := 1; i <= 5; i++ {
		numbers <- i
	}
}

func square(numbers <-chan int, squares chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(squares)

	for number := range numbers {
		squares <- number * number
	}
}

func printResult(squares <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for result := range squares {
		fmt.Println("Result:", result)
	}
}

func main() {
	numbers := make(chan int)
	squares := make(chan int)

	var wg sync.WaitGroup

	wg.Add(3)

	go generate(numbers, &wg)
	go square(numbers, squares, &wg)
	go printResult(squares, &wg)

	wg.Wait()

	fmt.Println("Pipeline completed")
}
