package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many goroutines should increment the counter? ")
	countText, _ := reader.ReadString('\n')
	count, err := strconv.Atoi(strings.TrimSpace(countText))
	if err != nil || count <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()

	fmt.Println("Final counter value:", atomic.LoadInt64(&counter))
}
