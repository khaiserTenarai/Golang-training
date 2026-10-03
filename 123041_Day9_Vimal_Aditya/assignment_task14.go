package main

import (
	"fmt"
	"sync"
	"time"
)

type File struct {
	ID   int
	Name string
	Size int
}

type Result struct {
	FileID   int
	FileName string
	Status   string
	Lines    int
}

func fileWorker(id int, jobs <-chan File, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for file := range jobs {
		fmt.Printf("   [Worker %d] Processing file: %s (%d KB)...\n", id, file.Name, file.Size)

		time.Sleep(time.Duration(file.Size*10) * time.Millisecond)

		linesProcessed := file.Size * 15

		results <- Result{
			FileID:   file.ID,
			FileName: file.Name,
			Status:   "SUCCESS",
			Lines:    linesProcessed,
		}

		fmt.Printf("   [Worker %d] Finished file: %s\n", id, file.Name)
	}
}

func main() {
	files := []File{
		{ID: 1, Name: "report_q1.csv", Size: 30},
		{ID: 2, Name: "user_logs.txt", Size: 50},
		{ID: 3, Name: "financials.json", Size: 20},
		{ID: 4, Name: "analytics.csv", Size: 40},
		{ID: 5, Name: "system_audit.log", Size: 10},
	}

	const numWorkers = 3
	jobs := make(chan File, len(files))
	results := make(chan Result, len(files))

	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go fileWorker(w, jobs, results, &wg)
	}

	for _, file := range files {
		jobs <- file
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Printf("%-8s %-20s %-10s %-15s\n", "FILE ID", "FILE NAME", "STATUS", "LINES PROCESSED")
	fmt.Println("-------------------------------------------------------")

	for res := range results {
		fmt.Printf("%-8d %-20s %-10s %-15d\n", res.FileID, res.FileName, res.Status, res.Lines)
	}

	fmt.Println("\n[Main] All files processed concurrently.")
}