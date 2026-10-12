package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many numbers do you want to put in the channel? ")
	countText, _ := reader.ReadString('\n')
	count, err := strconv.Atoi(strings.TrimSpace(countText))
	if err != nil || count < 0 {
		fmt.Println("Please enter a whole number, 0 or more.")
		return
	}

	ch := make(chan int, count)
	for i := 0; i < count; i++ {
		fmt.Printf("Number %d: ", i+1)
		numText, _ := reader.ReadString('\n')
		num, _ := strconv.Atoi(strings.TrimSpace(numText))
		ch <- num
	}
	close(ch)

	// Deliberately try to receive one extra time, to show what a
	// closed, empty channel looks like.
	for i := 0; i < count+1; i++ {
		value, ok := <-ch
		if !ok {
			fmt.Println("Channel is closed and empty now.")
			break
		}
		fmt.Println("Got:", value)
	}
}
