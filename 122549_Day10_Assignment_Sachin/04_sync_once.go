package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var once sync.Once

func setup() {
	fmt.Println("Running setup - this should only print once.")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many goroutines should try to run setup? ")
	countText, _ := reader.ReadString('\n')
	count, err := strconv.Atoi(strings.TrimSpace(countText))
	if err != nil || count <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	var wg sync.WaitGroup
	for i := 1; i <= count; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			once.Do(setup)
			fmt.Printf("Goroutine %d carried on after setup\n", id)
		}(i)
	}
	wg.Wait()
}
