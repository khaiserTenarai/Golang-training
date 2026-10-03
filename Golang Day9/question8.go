package main

import (
	"fmt"
	"time"
)

func main() {
	
	taskChan := make(chan string)

	go func() {
		tasks := []string{"Wash car", "Mow lawn", "Paint fence"}

		for _, task := range tasks {
			fmt.Printf("[Sender] Sending task: %s\n", task)
			taskChan <- task
			time.Sleep(500 * time.Millisecond)
		}

		fmt.Println("[Sender] All tasks sent! Closing the channel.")
		close(taskChan)
	}()

	fmt.Println("[Main] Waiting for tasks...")

	for task := range taskChan {
		fmt.Printf("[Main] Received and completed: %s\n", task)
	}

	fmt.Println("[Main] The channel is closed and all tasks are done. Exiting.")
}