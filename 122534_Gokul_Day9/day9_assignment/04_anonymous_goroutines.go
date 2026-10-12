package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Anonymous goroutine without any arguments")
	}()
	wg.Wait()

	names := []string{"Asha", "Ravi", "Meena"}
	for _, n := range names {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			fmt.Println("Welcome to the team,", name)
		}(n)
	}
	wg.Wait()
}
