package main

import (
	"fmt"
	"sync"
	"time"
)

var data int
var mu sync.RWMutex

func readData(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	mu.RLock()
	fmt.Println("Reader ", id, " read data: ", data)
	time.Sleep(500 * time.Millisecond)
	mu.RUnlock()
}
func writeData(wg *sync.WaitGroup, value int) {
	defer wg.Done()
	mu.Lock()
	data = value
	fmt.Println("Writer updated data to: ", data)
	mu.Unlock()
}
func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go readData(&wg, 1)
	go readData(&wg, 2)
	go writeData(&wg, 100)
	wg.Wait()
	fmt.Println("Final data: ", data)
}
