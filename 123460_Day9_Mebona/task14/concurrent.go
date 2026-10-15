package main

import (
	"fmt"
	"sync"
	"time"
)

type FileTask struct {
	FileName string
	SizeKB   int
}

func fileWorker(workerID int, tasks <-chan FileTask, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		// Simulate processing a file
		time.Sleep(40 * time.Millisecond)
		results <- fmt.Sprintf("Worker %d processed %s (%d KB)", workerID, task.FileName, task.SizeKB)
	}
}

func main() {
	tasks := make(chan FileTask, 5)
	results := make(chan string, 5)
	var wg sync.WaitGroup

	// Spawn workers
	numWorkers := 3
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go fileWorker(w, tasks, results, &wg)
	}

	// Enqueue files
	files := []FileTask{
		{"report_q1.csv", 450},
		{"employee_records.json", 1200},
		{"payroll_summary.pdf", 850},
		{"metrics.log", 300},
	}

	for _, file := range files {
		tasks <- file
	}
	close(tasks)

	// Wait for workers to finish in background, then close results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print results
	for res := range results {
		fmt.Println(res)
	}
}