package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var counter int

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many goroutines should increment the counter? ")
	countText, _ := reader.ReadString('\n')
	count, err := strconv.Atoi(strings.TrimSpace(countText))
	if err != nil || count <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // unprotected - this is the race
		}()
	}
	wg.Wait()

	fmt.Println("Final counter value:", counter, "(probably wrong - that's the bug)")
}
