// Tasks 7–10: Concurrency Patterns
// Task 7  — Worker pool
// Task 8  — Fan-out
// Task 9  — Fan-in
// Task 10 — Concurrent pipeline
// Run: go run main.go

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ============================================================
// Task 7: Build a Worker Pool
// ============================================================

type Job struct {
	ID   int
	Data int
}

type Result struct {
	JobID    int
	Output   int
	WorkerID int
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		time.Sleep(time.Duration(50+rand.Intn(100)) * time.Millisecond)
		results <- Result{
			JobID:    job.ID,
			Output:   job.Data * job.Data, // square the number
			WorkerID: id,
		}
	}
}

func demoWorkerPool() {
	fmt.Println("========== TASK 7: Worker Pool ==========")

	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Data: j * 5}
	}
	close(jobs)

	go func() { wg.Wait(); close(results) }()

	for r := range results {
		fmt.Printf("  Job #%d: %d² = %d (Worker %d)\n", r.JobID, int(0), r.Output, r.WorkerID)
	}
	fmt.Println()
}

// ============================================================
// Task 8: Implement Fan-Out
// ============================================================

func demoFanOut() {
	fmt.Println("========== TASK 8: Fan-Out ==========")

	// One input channel, distributed to multiple workers
	input := make(chan int, 10)
	numWorkers := 3
	var wg sync.WaitGroup

	// Fan-out: multiple goroutines read from the SAME channel
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for val := range input {
				time.Sleep(50 * time.Millisecond)
				fmt.Printf("  [Worker %d] Processed: %d → %d\n", id, val, val*2)
			}
		}(i)
	}

	// Send data
	for i := 1; i <= 9; i++ {
		input <- i * 10
	}
	close(input)

	wg.Wait()
	fmt.Println()
}

// ============================================================
// Task 9: Implement Fan-In
// ============================================================

func generate(name string, values ...int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, v := range values {
			time.Sleep(time.Duration(50+rand.Intn(100)) * time.Millisecond)
			out <- fmt.Sprintf("[%s] %d", name, v)
		}
	}()
	return out
}

// fanIn merges multiple channels into one
func fanIn(channels ...<-chan string) <-chan string {
	merged := make(chan string)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan string) {
			defer wg.Done()
			for val := range c {
				merged <- val
			}
		}(ch)
	}

	go func() { wg.Wait(); close(merged) }()
	return merged
}

func demoFanIn() {
	fmt.Println("========== TASK 9: Fan-In ==========")

	// Three separate producers
	ch1 := generate("API", 100, 200, 300)
	ch2 := generate("DB", 400, 500, 600)
	ch3 := generate("Cache", 700, 800, 900)

	// Fan-in: merge all into one channel
	merged := fanIn(ch1, ch2, ch3)

	for val := range merged {
		fmt.Printf("  [Fan-In] Received: %s\n", val)
	}
	fmt.Println()
}

// ============================================================
// Task 10: Build a Concurrent Pipeline
// ============================================================

// Stage 1: Generate numbers
func genNumbers(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

// Stage 2: Square numbers
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

// Stage 3: Double numbers
func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * 2
		}
	}()
	return out
}

// Stage 4: Filter (keep only > threshold)
func filter(in <-chan int, threshold int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n > threshold {
				out <- n
			}
		}
	}()
	return out
}

func demoPipeline() {
	fmt.Println("========== TASK 10: Concurrent Pipeline ==========")
	fmt.Println("  Pipeline: Generate → Square → Double → Filter(>100)")

	// Chain the pipeline stages
	stage1 := genNumbers(2, 3, 4, 5, 6, 7, 8, 9, 10)
	stage2 := square(stage1)
	stage3 := double(stage2)
	stage4 := filter(stage3, 100)

	for result := range stage4 {
		fmt.Printf("  Output: %d\n", result)
	}
	fmt.Println()
}

func main() {
	demoWorkerPool()
	demoFanOut()
	demoFanIn()
	demoPipeline()
}
