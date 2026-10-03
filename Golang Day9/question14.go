package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	
	defer wg.Done()

	for file := range jobs {
		fmt.Printf("[Worker %d] Started processing %s\n", id, file)
		
		time.Sleep(1 * time.Second)
		
		results <- fmt.Sprintf("%s completed by Worker %d", file, id)
	}
}

func main() {
	filesToProcess := []string{"sales.csv", "users.csv", "logs.txt", "config.json", "data.xml"}

	jobs := make(chan string, len(filesToProcess))
	results := make(chan string, len(filesToProcess))
	var wg sync.WaitGroup

	fmt.Println("Starting 3 workers...")
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for _, file := range filesToProcess {
		jobs <- file
	}
	
	close(jobs)

	go func() {
		wg.Wait() 
		close(results) 
	}()

	fmt.Println("Waiting for results...\n")

	
	for result := range results {
		fmt.Println("RESULT:", result)
	}

	fmt.Println("\nAll files processed successfully!")
}