package main

import (
	"fmt"
	"sync"
)

func producer(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	
	fmt.Println("[Producer] Generating data...")
	ch <- "Top Secret Data" 
	
}

func consumer(ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	
	fmt.Println("[Consumer] Waiting for data...")
	data := <-ch 
	
	fmt.Printf("[Consumer] Successfully processed: '%s'\n", data)
	
}

func main() {
	
	dataChan := make(chan string)
	var wg sync.WaitGroup

	wg.Add(2)

	go producer(dataChan, &wg)
	go consumer(dataChan, &wg)

	wg.Wait()
	fmt.Println("[Main] Program finished!")
}