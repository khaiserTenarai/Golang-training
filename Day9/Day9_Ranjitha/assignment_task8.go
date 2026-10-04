package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i * 100
		}
		close(ch)
		fmt.Println("[Producer] Channel closed.")
	}()

	fmt.Println("--- Approach 1: Reading via range ---")
	for value := range ch {
		fmt.Println("Received:", value)
	}

	fmt.Println("\n--- Approach 2: Comma-Ok Idiom after closing ---")
	
	closedCh := make(chan string, 1)
	closedCh <- "Active Data"
	close(closedCh)

	val, ok := <-closedCh
	fmt.Printf("Value: %s | Is Open: %t\n", val, ok)

	val, ok = <-closedCh
	fmt.Printf("Value: '%s' | Is Open: %t\n", val, ok)
}