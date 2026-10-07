package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting a task...\n", id)
	time.Sleep(time.Second) 
	fmt.Printf("Worker %d finished the task!\n", id)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Main: Starting the workers...")

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	fmt.Println("Main: Waiting for all workers to finish...")
	wg.Wait()
	fmt.Println("Main: All workers are done! The program can now exit safely.")
}