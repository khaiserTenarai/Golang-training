package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func processLine(line string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	upper := strings.ToUpper(line)
	results <- upper
}

func main() {
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var wg sync.WaitGroup
	results := make(chan string, 100)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1)
		go processLine(line, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for processed := range results {
		fmt.Println("Processed:", processed)
	}
}