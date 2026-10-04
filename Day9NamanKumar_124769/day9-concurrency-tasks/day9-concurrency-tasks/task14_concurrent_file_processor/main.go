package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

type FileResult struct {
	FileName  string
	WordCount int
}

func createSampleFiles(dir string, count int) []string {
	os.MkdirAll(dir, 0755)
	var files []string
	for i := 1; i <= count; i++ {
		name := fmt.Sprintf("%s/file%d.txt", dir, i)
		content := strings.Repeat(fmt.Sprintf("word%d ", i), i*3)
		os.WriteFile(name, []byte(content), 0644)
		files = append(files, name)
	}
	return files
}

func worker(id int, jobs <-chan string, results chan<- FileResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for file := range jobs {
		data, err := os.ReadFile(file)
		if err != nil {
			continue
		}
		words := strings.Fields(string(data))
		results <- FileResult{FileName: file, WordCount: len(words)}
	}
}

func main() {
	dir := "sample_files"
	files := createSampleFiles(dir, 5)

	jobs := make(chan string, len(files))
	results := make(chan FileResult, len(files))

	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for _, f := range files {
		jobs <- f
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Printf("%s: %d words\n", r.FileName, r.WordCount)
	}

	os.RemoveAll(dir)
}
