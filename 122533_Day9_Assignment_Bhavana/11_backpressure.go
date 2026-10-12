package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Buffer size (try a small number, like 2): ")
	bufText, _ := reader.ReadString('\n')
	bufSize, _ := strconv.Atoi(strings.TrimSpace(bufText))

	fmt.Print("How many items to produce: ")
	itemsText, _ := reader.ReadString('\n')
	items, _ := strconv.Atoi(strings.TrimSpace(itemsText))

	fmt.Print("Consumer delay per item, in seconds: ")
	delayText, _ := reader.ReadString('\n')
	delay, _ := strconv.Atoi(strings.TrimSpace(delayText))

	ch := make(chan int, bufSize)
	done := make(chan struct{})

	// slow consumer
	go func() {
		for item := range ch {
			time.Sleep(time.Duration(delay) * time.Second)
			fmt.Println("Consumed:", item)
		}
		close(done)
	}()

	// fast producer - once the buffer is full, this send blocks and
	// waits, which is backpressure in action.
	for i := 1; i <= items; i++ {
		fmt.Println("Producing:", i)
		ch <- i
	}
	close(ch)

	<-done
	fmt.Println("All items consumed.")
}
