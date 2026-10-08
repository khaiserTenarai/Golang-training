package main

import (
	"fmt"
	"sync"
	"time"
)

func workers(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Println("Worker", id, " started job ", job)
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Worker", id, " completed job ", job)
	}
}
func main() {
	jobs := make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go workers(i, jobs, &wg)
	}
	for i := 1; i <= 6; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	fmt.Println("All jobs completed")
}
