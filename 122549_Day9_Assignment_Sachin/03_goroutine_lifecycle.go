package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func namedWorker(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[named goroutine] started and finished")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many anonymous goroutines should we create? ")
	countText, _ := reader.ReadString('\n')
	count, err := strconv.Atoi(strings.TrimSpace(countText))
	if err != nil || count <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	var wg sync.WaitGroup

	// CREATED: this is the moment the named goroutine comes into existence.
	wg.Add(1)
	go namedWorker(&wg)

	for i := 1; i <= count; i++ {
		wg.Add(1)
		// an anonymous goroutine - a function with no name, defined
		// and launched right here, on the spot.
		go func(id int) {
			defer wg.Done()
			fmt.Printf("[anonymous goroutine %d] started and finished\n", id)
		}(i)
	}

	// Until wg.Wait() returns, every goroutine above is somewhere
	// between "created" and "terminated". Once Wait() returns, every
	// single one of them has reached the terminated stage.
	wg.Wait()
	fmt.Println("All goroutines have terminated.")
}
