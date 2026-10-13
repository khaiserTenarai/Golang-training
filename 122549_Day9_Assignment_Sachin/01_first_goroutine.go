package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func greet(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello,", name, "- this line was printed by a goroutine!")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What's your name? ")
	nameText, _ := reader.ReadString('\n')
	name := strings.TrimSpace(nameText)

	var wg sync.WaitGroup
	wg.Add(1)
	go greet(name, &wg)

	// main() would normally exit before the goroutine gets a chance
	// to run, so we wait here until it signals it's done.
	wg.Wait()
	fmt.Println("Back in main - the goroutine has finished.")
}
