package main

import (
	"fmt"
	"os"
	"sync"
)

type FileResult struct {
	FileName string
	Content  string
	Error    error
}

func processFile(fileName string, results chan<- FileResult, wg *sync.WaitGroup) {
	defer wg.Done()

	content, err := os.ReadFile(fileName)

	results <- FileResult{
		FileName: fileName,
		Content:  string(content),
		Error:    err,
	}
}

func main() {
	files := []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
	}

	results := make(chan FileResult)

	var wg sync.WaitGroup

	for _, fileName := range files {
		wg.Add(1)

		go processFile(fileName, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if result.Error != nil {
			fmt.Println("Error processing:", result.FileName)
			continue
		}

		fmt.Printf(
			"\nFile: %s\n%s",
			result.FileName,
			result.Content,
		)
	}

	fmt.Println("\nAll files processed")
}
