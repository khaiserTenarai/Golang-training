package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func square(n int, index int, results []int, wg *sync.WaitGroup) {
	defer wg.Done()
	results[index] = n * n
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many numbers do you want to square? ")
	countText, _ := reader.ReadString('\n')
	count, err := strconv.Atoi(strings.TrimSpace(countText))
	if err != nil || count <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Printf("Number %d: ", i+1)
		numText, _ := reader.ReadString('\n')
		num, _ := strconv.Atoi(strings.TrimSpace(numText))
		numbers[i] = num
	}

	results := make([]int, count)
	var wg sync.WaitGroup
	for i, n := range numbers {
		wg.Add(1)
		go square(n, i, results, &wg)
	}
	wg.Wait()

	fmt.Println("\nResults:")
	for i, n := range numbers {
		fmt.Printf("%d squared is %d\n", n, results[i])
	}
}
