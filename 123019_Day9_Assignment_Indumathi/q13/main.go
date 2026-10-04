package main

import (
	"fmt"
	"os"
	"sync"
)

type FileResult struct {
	FileName string
	Lines    int
	Words    int
	Chars    int
	Error    error
}

func processFile(fileName string, results chan<- FileResult, wg *sync.WaitGroup) {
	defer wg.Done()

	data, err := os.ReadFile(fileName)

	if err != nil {
		results <- FileResult{
			FileName: fileName,
			Error:    err,
		}
		return
	}

	content := string(data)

	lines := 0
	words := 0
	chars := len(content)

	for i := 0; i < len(content); i++ {
		if content[i] == '\n' {
			lines++
		}
	}

	if len(content) > 0 {
		lines++
	}

	inWord := false

	for i := 0; i < len(content); i++ {
		if content[i] != ' ' && content[i] != '\n' && content[i] != '\t' {
			if !inWord {
				words++
				inWord = true
			}
		} else {
			inWord = false
		}
	}

	results <- FileResult{
		FileName: fileName,
		Lines:    lines,
		Words:    words,
		Chars:    chars,
	}
}

func main() {
	files := []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
	}

	results := make(chan FileResult, len(files))

	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)

		go processFile(file, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if result.Error != nil {
			fmt.Println("Error processing", result.FileName, ":", result.Error)
			continue
		}

		fmt.Printf(
			"File: %s | Lines: %d | Words: %d | Characters: %d\n",
			result.FileName,
			result.Lines,
			result.Words,
			result.Chars,
		)
	}

	fmt.Println("All files processed successfully.")
}