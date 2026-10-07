package main

import (
	"fmt"
	"sync"
)

var data = 100
var mutex sync.RWMutex

func readData(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	mutex.RLock() // Allows multiple readers at the same time.
	fmt.Println("Reader", id, "read:", data)
	mutex.RUnlock()
}

func writeData(wg *sync.WaitGroup) {
	defer wg.Done()

	mutex.Lock() // Only one writer can change the data.
	data += 10
	fmt.Println("Writer changed data to:", data)
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(4)

	go readData(&wg, 1)
	go readData(&wg, 2)
	go writeData(&wg)
	go readData(&wg, 3)

	wg.Wait()
}
