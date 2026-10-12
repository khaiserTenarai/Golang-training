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

	fmt.Print("Delay in seconds for goroutine A: ")
	aText, _ := reader.ReadString('\n')
	aDelay, _ := strconv.Atoi(strings.TrimSpace(aText))

	fmt.Print("Delay in seconds for goroutine B: ")
	bText, _ := reader.ReadString('\n')
	bDelay, _ := strconv.Atoi(strings.TrimSpace(bText))

	chA := make(chan string)
	chB := make(chan string)

	go func() {
		time.Sleep(time.Duration(aDelay) * time.Second)
		chA <- "Message from A"
	}()
	go func() {
		time.Sleep(time.Duration(bDelay) * time.Second)
		chB <- "Message from B"
	}()

	// We're expecting exactly two messages, one from each goroutine,
	// so we select twice.
	for i := 0; i < 2; i++ {
		select {
		case msgA := <-chA:
			fmt.Println("Received:", msgA)
		case msgB := <-chB:
			fmt.Println("Received:", msgB)
		}
	}
}
