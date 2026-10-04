package day9sanskriti
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
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(fileName, "contains", len(data), "bytes")
}

func main() {

	files := []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
	}

	var wg sync.WaitGroup

	for _, file := range files {

		wg.Add(1)

		go processFile(file, &wg)
	}

	wg.Wait()

	fmt.Println("All files processed")
}