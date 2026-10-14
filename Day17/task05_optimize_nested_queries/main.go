package main

import (
	"fmt"
)

// Batch Loader / Dataloader pattern simulation for N+1 query optimization
type DataLoader struct {
	cache map[string]string
}

func (dl *DataLoader) LoadMany(ids []string) map[string]string {
	fmt.Printf("Batch fetching departments for IDs: %v (Single Query!)
", ids)
	return map[string]string{"d1": "Engineering", "d2": "HR"}
}

func main() {
	dl := &DataLoader{cache: make(map[string]string)}
	results := dl.LoadMany([]string{"d1", "d2"})
	fmt.Printf("Batch results: %+v
", results)
}
