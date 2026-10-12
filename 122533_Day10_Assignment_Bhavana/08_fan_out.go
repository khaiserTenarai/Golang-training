package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func fanOutWorker(id int, input <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range input {
		fmt.Printf("Goroutine %d picked up %d, doubled it to %d\n", id, n, n*2)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("How many numbers do you want to process? ")
	countText, _ := reader.ReadString('\n')
	count, err := strconv.Atoi(strings.TrimSpace(countText))
	if err != nil || count <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	fmt.Print("How many goroutines should share the work? ")
	workerText, _ := reader.ReadString('\n')
	workerCount, err := strconv.Atoi(strings.TrimSpace(workerText))
	if err != nil || workerCount <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	input := make(chan int, count)
	for i := 0; i < count; i++ {
		fmt.Printf("Number %d: ", i+1)
		numText, _ := reader.ReadString('\n')
		num, _ := strconv.Atoi(strings.TrimSpace(numText))
		input <- num
	}
	close(input)

	// this is the "fan-out" bit - several goroutines reading the one
	// input channel between them.
	var wg sync.WaitGroup
	for w := 1; w <= workerCount; w++ {
		wg.Add(1)
		go fanOutWorker(w, input, &wg)
	}
	wg.Wait()
}
