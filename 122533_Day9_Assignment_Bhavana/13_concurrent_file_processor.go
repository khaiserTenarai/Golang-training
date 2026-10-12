package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type fileResult struct {
	path  string
	lines int
	err   error
}

func processFile(path string, results chan<- fileResult, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(path)
	if err != nil {
		results <- fileResult{path: path, err: err}
		return
	}
	defer file.Close()

	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}

	results <- fileResult{path: path, lines: lineCount}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many files do you want to process? ")
	countText, _ := reader.ReadString('\n')
	count, err := strconv.Atoi(strings.TrimSpace(countText))
	if err != nil || count <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	paths := make([]string, count)
	for i := 0; i < count; i++ {
		fmt.Printf("File path %d: ", i+1)
		pathText, _ := reader.ReadString('\n')
		paths[i] = strings.TrimSpace(pathText)
	}

	results := make(chan fileResult, count)
	var wg sync.WaitGroup

	for _, path := range paths {
		wg.Add(1)
		go processFile(path, results, &wg)
	}

	// Close the results channel once every goroutine is done, so the
	// "range results" loop below knows when to stop instead of
	// hanging forever.
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if result.err != nil {
			fmt.Println(result.path, "- could not be read:", result.err)
			continue
		}
		fmt.Println(result.path, "-", result.lines, "lines")
	}
}
