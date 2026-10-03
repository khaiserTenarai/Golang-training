package main

import (
	"fmt"
	"os"
	"sync"
)

func processFile(fileName string, wg *sync.WaitGroup) {
	defer wg.Done()
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading ", fileName, ": ", err)
		return
	}
	fmt.Println("Processing:", fileName)
	fmt.Println(string(data))
}

func main() {

	files := []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
	}
	var wg sync.WaitGroup
	wg.Add(len(files))
	for _, file := range files {
		go processFile(file, &wg)
	}
	wg.Wait()
	fmt.Println("All files processed")
}
