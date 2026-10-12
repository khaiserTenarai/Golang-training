package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type job struct {
	id    int
	value int
}

type result struct {
	id      int
	value   int
	squared int
}

func worker(ctx context.Context, id int, jobs <-chan job, results chan<- result, processed *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		select {
		case <-ctx.Done():
			return // the timeout hit - stop picking up new jobs
		default:
		}

		time.Sleep(200 * time.Millisecond) // pretend this job takes a little while
		atomic.AddInt64(processed, 1)
		results <- result{id: j.id, value: j.value, squared: j.value * j.value}
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

	fmt.Print("Timeout, in seconds: ")
	timeoutText, _ := reader.ReadString('\n')
	timeoutSeconds, err := strconv.Atoi(strings.TrimSpace(timeoutText))
	if err != nil || timeoutSeconds <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	jobValues := make([]int, jobCount)
	for i := 0; i < jobCount; i++ {
		fmt.Printf("Value for job %d: ", i+1)
		numText, _ := reader.ReadString('\n')
		num, _ := strconv.Atoi(strings.TrimSpace(numText))
		jobValues[i] = num
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	jobs := make(chan job, jobCount)
	results := make(chan result, jobCount)
	var processed int64
	var wg sync.WaitGroup

	// fan-out: a pool of workers all pulling from the same jobs channel
	for w := 1; w <= workerCount; w++ {
		wg.Add(1)
		go worker(ctx, w, jobs, results, &processed, &wg)
	}

	for i, v := range jobValues {
		jobs <- job{id: i + 1, value: v}
	}
	close(jobs)

	// fan-in: once every worker is done, it's safe to close results
	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("\nCollecting results:")
collect:
	for {
		select {
		case res, ok := <-results:
			if !ok {
				break collect
			}
			fmt.Printf("Job %d: %d squared is %d\n", res.id, res.value, res.squared)
		case <-ctx.Done():
			fmt.Println("Timed out before every job finished.")
			break collect
		}
	}

	fmt.Println("\nJobs actually processed (race-safe count):", atomic.LoadInt64(&processed), "out of", jobCount)
}
