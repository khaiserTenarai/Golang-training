package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 3)

	fmt.Println("[Sender] Buffer capacity:", cap(ch))
	
	ch <- "Order-101"
	fmt.Println("[Sender] Sent: Order-101 (Buffer size:", len(ch), ")")

	ch <- "Order-102"
	fmt.Println("[Sender] Sent: Order-102 (Buffer size:", len(ch), ")")

	ch <- "Order-103"
	fmt.Println("[Sender] Sent: Order-103 (Buffer size:", len(ch), ")")

	go func() {
		fmt.Println("\n[Worker] Starting item processing...")
		time.Sleep(500 * time.Millisecond)

		fmt.Println("[Worker] Received:", <-ch)
		fmt.Println("[Worker] Received:", <-ch)
		fmt.Println("[Worker] Received:", <-ch)
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("\n[Main] Processing complete.")
}