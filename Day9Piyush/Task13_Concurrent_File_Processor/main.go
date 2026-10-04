// Task 13: Build a Concurrent File Processor
// Simulates processing multiple files concurrently using goroutines, channels,
// and a worker pool pattern.
// Run: go run main.go

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// FileJob represents a file to process
type FileJob struct {
	FileName string
	SizeKB   int
}

// FileResult represents the result of processing a file
type FileResult struct {
	FileName  string
	Status    string
	Duration  time.Duration
	WordCount int
	WorkerID  int
}

// worker processes files from the jobs channel
func worker(id int, jobs <-chan FileJob, results chan<- FileResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		start := time.Now()
		fmt.Printf("  [Worker %d] Processing: %s (%d KB)\n", id, job.FileName, job.SizeKB)

		// Simulate file processing (reading, parsing, counting)
		processingTime := time.Duration(job.SizeKB/10+rand.Intn(100)) * time.Millisecond
		time.Sleep(processingTime)

		wordCount := job.SizeKB * (10 + rand.Intn(20)) // simulated word count

		results <- FileResult{
			FileName:  job.FileName,
			Status:    "completed",
			Duration:  time.Since(start),
			WordCount: wordCount,
			WorkerID:  id,
		}
	}
}

func main() {
	fmt.Println("========== TASK 13: Concurrent File Processor ==========\n")

	// Simulated files to process
	files := []FileJob{
		{"report_2024.txt", 250},
		{"employees.csv", 180},
		{"salary_data.json", 320},
		{"attendance_log.txt", 150},
		{"performance_reviews.doc", 400},
		{"budget_plan.xlsx", 200},
		{"meeting_notes.txt", 90},
		{"project_timeline.csv", 280},
	}

	numWorkers := 3
	jobs := make(chan FileJob, len(files))
	results := make(chan FileResult, len(files))
	var wg sync.WaitGroup

	// Start worker pool
	fmt.Printf("Starting %d workers to process %d files...\n\n", numWorkers, len(files))
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send all jobs
	for _, file := range files {
		jobs <- file
	}
	close(jobs)

	// Wait for workers and close results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect and display results
	fmt.Println("\n--- Processing Results ---")
	totalWords := 0
	totalDuration := time.Duration(0)
	for result := range results {
		fmt.Printf("  %-30s | Worker %d | %v | %d words\n",
			result.FileName, result.WorkerID, result.Duration.Round(time.Millisecond), result.WordCount)
		totalWords += result.WordCount
		totalDuration += result.Duration
	}

	fmt.Printf("\n--- Summary ---\n")
	fmt.Printf("  Files processed: %d\n", len(files))
	fmt.Printf("  Total words:     %d\n", totalWords)
	fmt.Printf("  Workers used:    %d\n", numWorkers)
	fmt.Println("\nDone!")
}
