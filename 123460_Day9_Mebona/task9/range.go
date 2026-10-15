package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Alice"
		ch <- "Bob"
		ch <- "Charlie"
		close(ch) // Crucial: loop will deadlock without closing
	}()

	// Range automatically breaks when the channel is closed and emptied
	for name := range ch {
		fmt.Printf("Processed employee: %s\n", name)
	}
}