package main

import (
	"fmt"
	"os"
	"sync"
)

/*
	CONCURRENT FILE PROCESSOR

	Producer / Main
	      |
	      +---- file1.txt
	      +---- file2.txt
	      +---- file3.txt
	      |
	      ↓
	Multiple Goroutines
	      |
	      ↓
	Process files concurrently


	FILE CREATION NOTES:

	1. os.Create("file.txt")

	   - If file does NOT exist → creates the file.
	   - If file ALREADY exists → clears/truncates its old content.
	   - Then we can write new content.

	2. os.WriteFile("file.txt", data, 0644)

	   - If file does NOT exist → creates the file.
	   - If file ALREADY exists → replaces its old content.

	3. os.OpenFile() with O_CREATE | O_EXCL

	   - If file does NOT exist → creates the file.
	   - If file ALREADY exists → returns an error.
	   - Useful when we want "create only if it does not exist".

	4. os.OpenFile() with O_CREATE | O_APPEND

	   - If file does NOT exist → creates the file.
	   - If file ALREADY exists → keeps old content and adds
	     new content at the end.

	In this example, os.WriteFile() is used only to create
	sample input files for demonstration.
*/

func processFile(fileName string, wg *sync.WaitGroup) {

	// Tell WaitGroup that this goroutine is finished.
	defer wg.Done()

	fmt.Println("Processing:", fileName)

	/*
		Read the existing file.

		We are NOT creating the file here.
		The file is already created before processing.
	*/
	data, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(
		"Completed:",
		fileName,
		"| Size:",
		len(data),
		"bytes",
	)
}

func main() {

	/*
		CREATE SAMPLE FILES

		os.WriteFile() is used here.

		If the file does not exist:
			→ It creates the file.

		If the file already exists:
			→ It replaces the existing content.

		For a real file processor, we would normally
		process existing files instead of creating them.
	*/

	os.WriteFile("file1.txt", []byte("Hello Go"), 0644)

	os.WriteFile("file2.txt", []byte("Go Programming"), 0644)

	os.WriteFile("file3.txt", []byte("Concurrency"), 0644)

	files := []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
	}

	var wg sync.WaitGroup

	/*
		Process each file concurrently.

		"go" creates a goroutine.

		Each file gets its own goroutine.
	*/

	for _, fileName := range files {

		wg.Add(1)

		go processFile(fileName, &wg)
	}

	/*
		Wait until all goroutines finish.
	*/
	wg.Wait()

	fmt.Println("All files processed")
}

/*
Output:
-----------
PS C:\Training\HP - GO Language Training\Day_9\124772_Day_9(Go)_Assignment_Reddem_Ganesh_Reddy\13-Process Files Concurrently> go run .\main.go
Processing: file3.txt
Processing: file2.txt
Processing: file1.txt
Completed: file3.txt | Size: 11 bytes
Completed: file2.txt | Size: 14 bytes
Completed: file1.txt | Size: 8 bytes
All files processed
*/