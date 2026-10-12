package main

import (
	"fmt"
	"time"
)

func lifeCycleWorker(start <-chan bool, done chan<- bool) {
	fmt.Println("2. Goroutine is running, but now it is blocked waiting for a signal")
	<-start
	fmt.Println("4. Signal received, goroutine is working again")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("5. Goroutine has finished its work")
	done <- true
}

func main() {
	start := make(chan bool)
	done := make(chan bool)

	fmt.Println("1. Creating the goroutine")
	go lifeCycleWorker(start, done)

	time.Sleep(1 * time.Second)
	fmt.Println("3. Main is sending the start signal")
	start <- true

	<-done
	fmt.Println("6. Goroutine has exited, main is ending")
}
