package main

import (
	"fmt"
	"time"
)

func main() {

	jobChannel := make(chan string)

	go func() {
		jobs := []string{"Email Notification", "Invoice Generation", "Data Backup", "Report Export"}

		for _, job := range jobs {
			fmt.Println("[Producer] Sending job:", job)
			jobChannel <- job
			time.Sleep(200 * time.Millisecond)
		}

		close(jobChannel)
		fmt.Println("[Producer] All jobs sent. Channel closed.")
	}()

	fmt.Println("[Consumer] Ready to process jobs...")
	for job := range jobChannel {
		fmt.Printf("[Consumer] Processed: %s\n", job)
	}

	fmt.Println("[Main] All jobs consumed cleanly via range. Program finished.")
}