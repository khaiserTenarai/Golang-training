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
	fmt.Print("How many words do you want to send? ")
	countText, _ := reader.ReadString('\n')
	count, err := strconv.Atoi(strings.TrimSpace(countText))
	if err != nil || count <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	ch := make(chan string, count)
	for i := 0; i < count; i++ {
		fmt.Printf("Word %d: ", i+1)
		wordText, _ := reader.ReadString('\n')
		ch <- strings.TrimSpace(wordText)
	}
	close(ch) // without this, "range ch" below would block forever

	fmt.Println("Reading back with range:")
	for word := range ch {
		fmt.Println("-", word)
	}
}
