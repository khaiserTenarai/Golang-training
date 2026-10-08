package main

import (
	"fmt"
	"sync"
)

var data1 sync.Map

func storeData(wg *sync.WaitGroup, key string, value string) {
	defer wg.Done()
	data1.Store(key, value)
	fmt.Println("Stored: ", key, value)
}
func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go storeData(&wg, "name", "Tom")
	go storeData(&wg, "skill", "Go Lang")
	go storeData(&wg, "company", "xyz")
	wg.Wait()
	fmt.Println("Reading data: ")
	data1.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
