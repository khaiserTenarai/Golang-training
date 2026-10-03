package main

import (
	"fmt"
	"sync"
	"time"
)

func lifecycleDemo(signalChan chan bool, wg *sync.WaitGroup) {
	// 5. TERMINATION: When the function ends, wg.Done() is called and the goroutine dies
	defer wg.Done()

	// 2. RUNNING: The goroutine has been assigned to a processor and is actively executing code
	fmt.Println("[Goroutine] State: RUNNING (Executing code)")

	// 3. BLOCKED: The goroutine stops and waits for something to happen.
	// In this case, it waits to receive a signal from the channel.
	fmt.Println("[Goroutine] State: BLOCKED (Waiting for a signal...)")
	<-signalChan 

	// 4. RUNNING AGAIN: The signal was received, so the goroutine wakes up
	fmt.Println("[Goroutine] State: RUNNING AGAIN (Signal received!)")
	fmt.Println("[Goroutine] State: TERMINATING (Function is ending)")
}

func main() {
	var wg sync.WaitGroup
	
	// Create a channel to send a signal between goroutines
	signalChan := make(chan bool)

	// 1. CREATION: The 'go' keyword creates the goroutine and puts it in the runtime's waiting queue
	fmt.Println("[Main] State: CREATION (Spawning new goroutine)")
	wg.Add(1)
	go lifecycleDemo(signalChan, &wg)

	// Pause main for a second to let the goroutine reach its BLOCKED state
	time.Sleep(1 * time.Second)

	// Send a signal (the value 'true') into the channel to wake the goroutine up
	fmt.Println("[Main] Unblocking the goroutine now...")
	signalChan <- true

	// Wait for the goroutine to finish its termination process
	wg.Wait()
	fmt.Println("[Main] The goroutine is dead. Program exiting.")
}