package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range jobs {
		fmt.Printf("Worker %d is processing %d\n", id, n)
		results <- n * n
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("How many jobs? ")
	jobCountText, _ := reader.ReadString('\n')
	jobCount, err := strconv.Atoi(strings.TrimSpace(jobCountText))
	if err != nil || jobCount <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	fmt.Print("How many workers? ")
	workerCountText, _ := reader.ReadString('\n')
	workerCount, err := strconv.Atoi(strings.TrimSpace(workerCountText))
	if err != nil || workerCount <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	jobs := make(chan int, jobCount)
	results := make(chan int, jobCount)
	var wg sync.WaitGroup

	for w := 1; w <= workerCount; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for i := 0; i < jobCount; i++ {
		fmt.Printf("Job %d value: ", i+1)
		numText, _ := reader.ReadString('\n')
		num, _ := strconv.Atoi(strings.TrimSpace(numText))
		jobs <- num
	}
	close(jobs)

	wg.Wait()
	close(results)

	fmt.Println("\nResults:")
	for result := range results {
		fmt.Println(result)
	}
}
