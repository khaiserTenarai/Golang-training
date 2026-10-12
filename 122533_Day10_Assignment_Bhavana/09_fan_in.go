package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func source(id int, howMany int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 1; i <= howMany; i++ {
			out <- fmt.Sprintf("message %d from source %d", i, id)
		}
	}()
	return out
}

// merge fans multiple input channels into one output channel.
func merge(channels ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan string) {
			defer wg.Done()
			for msg := range c {
				out <- msg
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("How many source goroutines? ")
	sourceCountText, _ := reader.ReadString('\n')
	sourceCount, err := strconv.Atoi(strings.TrimSpace(sourceCountText))
	if err != nil || sourceCount <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	fmt.Print("How many messages should each source send? ")
	perSourceText, _ := reader.ReadString('\n')
	perSource, err := strconv.Atoi(strings.TrimSpace(perSourceText))
	if err != nil || perSource <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	channels := make([]<-chan string, sourceCount)
	for i := 0; i < sourceCount; i++ {
		channels[i] = source(i+1, perSource)
	}

	for msg := range merge(channels...) {
		fmt.Println(msg)
	}
}
