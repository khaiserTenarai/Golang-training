package main
// Fan-out means one input source sends work to multiple goroutines.
import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("Worker", id, "processed", job)
	}
}

func main() {
	

	jobs := make(chan int)

	var wg sync.WaitGroup

	// Create 3 workers.
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Send jobs.
	for i := 1; i <= 6; i++ {
		jobs <- i
	}

	close(jobs)

	wg.Wait()

	fmt.Println("Fan-out completed")
}