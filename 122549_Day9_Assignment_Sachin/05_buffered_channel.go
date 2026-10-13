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

	fmt.Print("Buffer size: ")
	sizeText, _ := reader.ReadString('\n')
	size, err := strconv.Atoi(strings.TrimSpace(sizeText))
	if err != nil || size <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	ch := make(chan int, size)

	fmt.Printf("Enter %d numbers, one per line:\n", size)
	for i := 0; i < size; i++ {
		fmt.Printf("Number %d: ", i+1)
		numText, _ := reader.ReadString('\n')
		num, err := strconv.Atoi(strings.TrimSpace(numText))
		if err != nil {
			fmt.Println("That wasn't a number, skipping it.")
			continue
		}
		ch <- num // doesn't block - there's still room in the buffer
		fmt.Println("Sent", num, "- didn't block because the buffer isn't full yet")
	}

	close(ch)

	fmt.Println("\nNow reading everything back out:")
	for num := range ch {
		fmt.Println("Received:", num)
	}
}
