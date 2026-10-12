package main

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello from my first goroutine")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go sayHello(&wg)

	fmt.Println("Hello from main")
	wg.Wait()
}
