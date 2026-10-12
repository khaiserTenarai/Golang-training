package main

import "fmt"

func main() {
	ch := make(chan string, 3)

	ch <- "Asha"
	ch <- "Ravi"
	close(ch)

	name, ok := <-ch
	fmt.Printf("value: %q, open: %v\n", name, ok)

	name, ok = <-ch
	fmt.Printf("value: %q, open: %v\n", name, ok)

	name, ok = <-ch
	fmt.Printf("value: %q, open: %v\n", name, ok)

	// sending on a closed channel would cause a panic
}
