// Tasks 11, 12: Backpressure and Producer-Consumer Pattern
// Task 11 — Demonstrate backpressure
// Task 12 — Create producer-consumer
// Run: go run main.go

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ============================================================
// Task 11: Demonstrate Backpressure
// ============================================================

func demoBackpressure() {
	fmt.Println("========== TASK 11: Backpressure ==========")

	// Small buffer = backpressure when consumer is slow
	ch := make(chan int, 3) // only 3 slots

	// Fast producer
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Printf("  [Producer] Trying to send %d (buffer: %d/%d)...\n", i, len(ch), cap(ch))
			ch <- i // BLOCKS when buffer is full — this IS backpressure
			fmt.Printf("  [Producer] Sent %d\n", i)
		}
		close(ch)
	}()

	// Slow consumer
	for val := range ch {
		fmt.Printf("  [Consumer] Processing %d...\n", val)
		time.Sleep(300 * time.Millisecond) // simulate slow processing
	}

	fmt.Println("  Backpressure demo complete!\n")
}

// ============================================================
// Task 12: Producer-Consumer Pattern
// ============================================================

type Task struct {
	ID   int
	Data string
}

func producer(id int, tasks chan<- Task, count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		task := Task{
			ID:   id*100 + i + 1,
			Data: fmt.Sprintf("Task from producer %d", id),
		}
		tasks <- task
		fmt.Printf("  [Producer %d] Created task #%d\n", id, task.ID)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func consumer(id int, tasks <-chan Task, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("  [Consumer %d] Processing task #%d\n", id, task.ID)
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		results <- fmt.Sprintf("Task #%d completed by consumer %d", task.ID, id)
	}
}

func demoProducerConsumer() {
	fmt.Println("========== TASK 12: Producer-Consumer ==========")

	tasks := make(chan Task, 5)
	results := make(chan string, 10)

	var producerWg sync.WaitGroup
	var consumerWg sync.WaitGroup

	// Start 2 producers (each produces 3 tasks)
	for i := 1; i <= 2; i++ {
		producerWg.Add(1)
		go producer(i, tasks, 3, &producerWg)
	}

	// Start 3 consumers
	for i := 1; i <= 3; i++ {
		consumerWg.Add(1)
		go consumer(i, tasks, results, &consumerWg)
	}

	// Close tasks channel when all producers are done
	go func() {
		producerWg.Wait()
		close(tasks)
		fmt.Println("  [System] All producers finished, tasks channel closed.")
	}()

	// Collect results
	go func() {
		consumerWg.Wait()
		close(results)
		fmt.Println("  [System] All consumers finished, results channel closed.")
	}()

	// Print results
	fmt.Println("\n  --- Results ---")
	for result := range results {
		fmt.Printf("  ✓ %s\n", result)
	}

	fmt.Println("\n  Producer-Consumer demo complete!\n")
}

func main() {
	demoBackpressure()
	demoProducerConsumer()
}
