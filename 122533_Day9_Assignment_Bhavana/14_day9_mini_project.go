package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type employee struct {
	name   string
	salary float64
}

type bonusResult struct {
	name   string
	salary float64
	bonus  float64
}

// producer sends every employee into the jobs channel, then closes it
// so the workers know there's nothing more coming.
func producer(employees []employee, jobs chan<- employee) {
	for _, e := range employees {
		jobs <- e
	}
	close(jobs)
}

// worker is one of several goroutines pulling jobs off the same
// channel and doing the actual calculation.
func worker(id int, jobs <-chan employee, results chan<- bonusResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		bonus := job.salary * 0.10
		fmt.Printf("[worker %d] processed %s\n", id, job.name)
		results <- bonusResult{name: job.name, salary: job.salary, bonus: bonus}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("How many employees? ")
	countText, _ := reader.ReadString('\n')
	count, err := strconv.Atoi(strings.TrimSpace(countText))
	if err != nil || count <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	employees := make([]employee, count)
	for i := 0; i < count; i++ {
		fmt.Printf("\nEmployee %d name: ", i+1)
		nameText, _ := reader.ReadString('\n')

		fmt.Printf("Employee %d salary: ", i+1)
		salaryText, _ := reader.ReadString('\n')
		salary, err := strconv.ParseFloat(strings.TrimSpace(salaryText), 64)
		if err != nil {
			fmt.Println("That wasn't a number, using 0.")
			salary = 0
		}

		employees[i] = employee{name: strings.TrimSpace(nameText), salary: salary}
	}

	fmt.Print("\nHow many worker goroutines should process them? ")
	workerText, _ := reader.ReadString('\n')
	workerCount, err := strconv.Atoi(strings.TrimSpace(workerText))
	if err != nil || workerCount <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	jobs := make(chan employee, count)
	results := make(chan bonusResult, count)
	var wg sync.WaitGroup

	go producer(employees, jobs)

	for w := 1; w <= workerCount; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Once every worker is done, close the results channel - that's
	// what lets the collection loop below finish instead of hanging.
	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("\nCollecting results:")
	for {
		select {
		case result, ok := <-results:
			if !ok {
				fmt.Println("\nAll employees processed.")
				return
			}
			fmt.Printf("%s - salary %.2f -> bonus %.2f\n", result.name, result.salary, result.bonus)
		case <-time.After(5 * time.Second):
			fmt.Println("Timed out waiting for the workers - something's stuck.")
			return
		}
	}
}
