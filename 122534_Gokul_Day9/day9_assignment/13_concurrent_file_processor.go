/*
Question 13: Build a concurrent file processor.

Each file holds lines in the form "name,salary". Workers read the files
at the same time and report the employee count and salary total per file.
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type fileResult struct {
	name  string
	count int
	total int
	err   error
}

func processFile(path string) fileResult {
	res := fileResult{name: filepath.Base(path)}

	data, err := os.ReadFile(path)
	if err != nil {
		res.err = err
		return res
	}

	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		salary, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			continue
		}
		res.count++
		res.total += salary
	}
	return res
}

func main() {
	dir, err := os.MkdirTemp("", "employees")
	if err != nil {
		fmt.Println("could not create temp folder:", err)
		return
	}
	defer os.RemoveAll(dir)

	files := map[string]string{
		"sales.txt":   "Asha,30000\nRavi,35000\nMeena,32000",
		"support.txt": "Karthik,28000\nDivya,27000",
		"tech.txt":    "Arun,50000\nSneha,52000\nVikram,48000\nLatha,55000",
	}

	var paths []string
	for name, text := range files {
		p := filepath.Join(dir, name)
		if err := os.WriteFile(p, []byte(text), 0644); err != nil {
			fmt.Println("could not write file:", err)
			return
		}
		paths = append(paths, p)
	}

	jobs := make(chan string)
	results := make(chan fileResult)
	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for p := range jobs {
				results <- processFile(p)
			}
		}()
	}

	go func() {
		for _, p := range paths {
			jobs <- p
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	var all []fileResult
	for r := range results {
		all = append(all, r)
	}
	sort.Slice(all, func(i, j int) bool { return all[i].name < all[j].name })

	for _, r := range all {
		if r.err != nil {
			fmt.Println(r.name, "failed:", r.err)
			continue
		}
		fmt.Printf("%-12s employees: %d, total salary: %d\n", r.name, r.count, r.total)
	}
}
