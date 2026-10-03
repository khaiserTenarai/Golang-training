package main

import (
	"fmt"
	"time"
)

func main() {

	queue := make(chan string, 3)

	go func() {
		for i := 1; i <= 6; i++ {
			task := fmt.Sprintf("Task %d", i)
			fmt.Printf("[Producer] Attempting to send %s...\n", task)
			
			queue <- task 
			
			fmt.Printf("[Producer] --> Successfully sent %s!\n", task)
		}
		close(queue)
		fmt.Println("[Producer] Finished sending all tasks!")
	}()

	
	time.Sleep(1 * time.Second)
	fmt.Println("\n[Consumer] Waking up! Starting to clear the queue...\n")

	
	for task := range queue {
		fmt.Printf("[Consumer] Received %s. Processing...\n", task)
		
		
		time.Sleep(2 * time.Second)
	}

	fmt.Println("All processing complete. Program exiting.")
}