package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func processFile(filename string) {
	fmt.Println("Processing:", filename)

	// Simulate file processing
	time.Sleep(1 * time.Second)

	fmt.Println("Completed:", filename)
}

func worker(id int, files <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for filename := range files {
		fmt.Printf("Worker %d started %s\n", id, filename)

		processFile(filename)

		fmt.Printf("Worker %d finished %s\n", id, filename)
	}
}

func main() {
	files := []string{
		"file1.txt",
		"file2.txt",
	}

	// Number of concurrent workers
	workerCount := 3

	jobs := make(chan string)

	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Send files to workers
	for _, file := range files {
		_, err := os.Stat(file)
		if err != nil {
			fmt.Println("File not found:", file)
			continue
		}

		jobs <- file
	}

	close(jobs)

	// Wait for all workers
	wg.Wait()

	fmt.Println("All files processed")
}
