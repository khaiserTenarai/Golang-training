package main

import (
	"fmt"
	"time"
)

func lifecycleDemo(id int) {
	fmt.Printf("[Created & Running] Goroutine %d started\n", id)
	
	// Simulating work / blocked state
	time.Sleep(200 * time.Millisecond)
	
	fmt.Printf("[Terminated] Goroutine %d finished\n", id)
}

func main() {
	fmt.Println("[Main] Starting goroutines...")
	for i := 1; i <= 2; i++ {
		go lifecycleDemo(i)
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println("[Main] Exiting...")
}