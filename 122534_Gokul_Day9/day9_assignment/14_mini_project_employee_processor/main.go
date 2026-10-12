/*
Question 14: Day 9 Mini Project - Build a concurrent employee processing
system using goroutines, channels, select, a producer, a consumer and workers.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	jobs := make(chan Employee, 3)
	results := make(chan Payslip)
	done := make(chan Summary, 1)

	var wg sync.WaitGroup

	go producer(jobs)

	for id := 1; id <= 3; id++ {
		wg.Add(1)
		go worker(id, jobs, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	go consumer(results, done)

	select {
	case summary := <-done:
		fmt.Printf("\nProcessed %d employees, total net payout: %.2f\n",
			summary.Count, summary.TotalNet)
	case <-time.After(10 * time.Second):
		fmt.Println("\nTimed out while processing employees")
	}
}
