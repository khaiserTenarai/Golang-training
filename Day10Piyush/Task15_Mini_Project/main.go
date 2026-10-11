// Task 15: Day 10 Mini Project — Cancellation-Aware Worker System
// Features: Worker pool, Fan-out, Fan-in, Timeout, Race-safe processing
// Run: go run main.go
// Run with race detector: go run -race main.go

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// ============================================================
// Models
// ============================================================

type Task struct {
	ID       int
	Type     string // "salary", "report", "email", "audit"
	Payload  string
	Priority int
}

type TaskResult struct {
	TaskID    int
	Type      string
	Output    string
	WorkerID  int
	Duration  time.Duration
	Status    string // "completed", "cancelled", "timeout"
}

// ============================================================
// Worker: Processes tasks, respects context cancellation
// ============================================================

func processTask(ctx context.Context, task Task, workerID int) TaskResult {
	start := time.Now()

	// Simulate variable processing time based on task type
	delay := time.Duration(100+rand.Intn(300)) * time.Millisecond

	select {
	case <-time.After(delay):
		// Task completed successfully
		output := ""
		switch task.Type {
		case "salary":
			output = fmt.Sprintf("Calculated salary for %s", task.Payload)
		case "report":
			output = fmt.Sprintf("Generated report: %s", task.Payload)
		case "email":
			output = fmt.Sprintf("Sent email to %s", task.Payload)
		case "audit":
			output = fmt.Sprintf("Audit completed for %s", task.Payload)
		}
		return TaskResult{
			TaskID: task.ID, Type: task.Type, Output: output,
			WorkerID: workerID, Duration: time.Since(start), Status: "completed",
		}
	case <-ctx.Done():
		return TaskResult{
			TaskID: task.ID, Type: task.Type, Output: ctx.Err().Error(),
			WorkerID: workerID, Duration: time.Since(start), Status: "cancelled",
		}
	}
}

// ============================================================
// Worker Pool with Fan-Out
// ============================================================

func startWorkerPool(ctx context.Context, numWorkers int, tasks <-chan Task, results chan<- TaskResult, wg *sync.WaitGroup, processed *int64) {
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for task := range tasks {
				// Check cancellation before processing
				select {
				case <-ctx.Done():
					results <- TaskResult{
						TaskID: task.ID, Type: task.Type,
						Output: "skipped — context cancelled", WorkerID: workerID,
						Duration: 0, Status: "cancelled",
					}
					continue
				default:
				}

				result := processTask(ctx, task, workerID)
				atomic.AddInt64(processed, 1) // Race-safe counter
				results <- result
			}
		}(i)
	}
}

// ============================================================
// Fan-In: Merge results from multiple worker types
// ============================================================

func fanIn(channels ...<-chan TaskResult) <-chan TaskResult {
	merged := make(chan TaskResult, 50)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan TaskResult) {
			defer wg.Done()
			for r := range c {
				merged <- r
			}
		}(ch)
	}

	go func() { wg.Wait(); close(merged) }()
	return merged
}

// ============================================================
// Report Generator
// ============================================================

func printReport(results []TaskResult, totalTasks int, elapsed time.Duration) {
	fmt.Println("\n╔════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                    TASK PROCESSING REPORT                             ║")
	fmt.Println("╠════════════════════════════════════════════════════════════════════════╣")
	fmt.Printf("║ %-5s %-10s %-35s %-8s %-6s ║\n", "Task", "Type", "Output", "Status", "Worker")
	fmt.Println("╠════════════════════════════════════════════════════════════════════════╣")

	completed, cancelled := 0, 0
	for _, r := range results {
		status := r.Status
		if status == "completed" {
			completed++
		} else {
			cancelled++
		}
		output := r.Output
		if len(output) > 35 {
			output = output[:32] + "..."
		}
		fmt.Printf("║ %-5d %-10s %-35s %-8s %-6d ║\n",
			r.TaskID, r.Type, output, status, r.WorkerID)
	}

	fmt.Println("╠════════════════════════════════════════════════════════════════════════╣")
	fmt.Printf("║ Total: %-3d | Completed: %-3d | Cancelled: %-3d | Time: %-14v ║\n",
		totalTasks, completed, cancelled, elapsed.Round(time.Millisecond))
	fmt.Println("╚════════════════════════════════════════════════════════════════════════╝")
}

// ============================================================
// Main — Orchestrates the System
// ============================================================

func main() {
	fmt.Println("===== DAY 10 MINI PROJECT: Cancellation-Aware Worker System =====\n")

	// Task data
	tasks := []Task{
		{1, "salary", "Piyush", 1},
		{2, "report", "Q4 Revenue", 2},
		{3, "email", "alice@test.com", 1},
		{4, "audit", "Finance Department", 3},
		{5, "salary", "Bob", 1},
		{6, "report", "Annual Summary", 2},
		{7, "email", "charlie@test.com", 1},
		{8, "audit", "Engineering Team", 3},
		{9, "salary", "Diana", 1},
		{10, "report", "Monthly KPIs", 2},
		{11, "email", "eve@test.com", 1},
		{12, "audit", "HR Compliance", 3},
		{13, "salary", "Frank", 1},
		{14, "report", "Budget Forecast", 2},
		{15, "email", "grace@test.com", 1},
	}

	start := time.Now()

	// Context with timeout — cancels all work after 1.5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()

	// Channels
	priorityTasks := make(chan Task, 10)   // high priority queue
	regularTasks := make(chan Task, 10)     // regular queue
	priorityResults := make(chan TaskResult, 20)
	regularResults := make(chan TaskResult, 20)

	var processed int64

	// Fan-Out: Two worker pools (priority + regular)
	var wg1, wg2 sync.WaitGroup
	fmt.Println("[System] Starting priority workers (2) and regular workers (3)...\n")

	startWorkerPool(ctx, 2, priorityTasks, priorityResults, &wg1, &processed)
	startWorkerPool(ctx, 3, regularTasks, regularResults, &wg2, &processed)

	// Distribute tasks by priority (fan-out)
	go func() {
		for _, task := range tasks {
			select {
			case <-ctx.Done():
				break
			default:
			}
			if task.Priority <= 1 {
				priorityTasks <- task
			} else {
				regularTasks <- task
			}
		}
		close(priorityTasks)
		close(regularTasks)
	}()

	// Close result channels when workers finish
	go func() { wg1.Wait(); close(priorityResults) }()
	go func() { wg2.Wait(); close(regularResults) }()

	// Fan-In: Merge results from both pools
	allResults := fanIn(priorityResults, regularResults)

	// Collect results
	var collected []TaskResult
	for r := range allResults {
		fmt.Printf("  [Result] Task #%-2d (%s) → %s by Worker %d [%v]\n",
			r.TaskID, r.Type, r.Status, r.WorkerID, r.Duration.Round(time.Millisecond))
		collected = append(collected, r)
	}

	elapsed := time.Since(start)
	printReport(collected, len(tasks), elapsed)

	fmt.Printf("\nAtomic processed count: %d\n", atomic.LoadInt64(&processed))
	fmt.Println("Run with 'go run -race main.go' to verify race safety.")
	fmt.Println("\nDone!")
}
